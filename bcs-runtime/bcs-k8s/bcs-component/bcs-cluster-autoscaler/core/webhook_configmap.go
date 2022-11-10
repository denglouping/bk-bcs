/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"encoding/json"
	"fmt"
	"strings"

	contextinternal "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cluster-autoscaler/context"
	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/autoscaler/cluster-autoscaler/clusterstate"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	scheduler_util "k8s.io/autoscaler/cluster-autoscaler/utils/scheduler"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

var _ Webhook = &ConfigMapScaler{}

// ConfigMapScaler impletements Webhook via configmap
type ConfigMapScaler struct {
	client    kubeclient.Interface
	namespace string
	name      string
}

// NewConfigMapScaler initilizes a ConfigMapScaler
func NewConfigMapScaler(client kubeclient.Interface, configmap string) Webhook {
	slice := strings.Split(configmap, "/")
	return &ConfigMapScaler{client: client, namespace: slice[0], name: slice[1]}
}

// DoWebhook get responses from webhook, then execute scale based on responses
func (c *ConfigMapScaler) DoWebhook(context *contextinternal.Context,
	clusterStateRegistry *clusterstate.ClusterStateRegistry, sd *ScaleDown,
	nodes []*corev1.Node, pods []*corev1.Pod) errors.AutoscalerError {
	nodeNameToNodeInfo := scheduler_util.CreateNodeNameToInfoMap(pods, nodes)

	options, candidates, err := c.GetResponses(context, clusterStateRegistry,
		nodeNameToNodeInfo, nodes, sd)
	if err != nil {
		return errors.NewAutoscalerError(errors.TransientError,
			"failed to get response from configmap: %v", err)
	}
	err = c.ExecuteScale(context, clusterStateRegistry, sd, nodes, options, candidates, nodeNameToNodeInfo)
	if err != nil {
		return errors.NewAutoscalerError(errors.ApiCallError,
			"failed to execute scale from configmap: %v", err)
	}
	return nil
}

// GetResponses returns the responses of webhook
func (c *ConfigMapScaler) GetResponses(context *contextinternal.Context,
	clusterStateRegistry *clusterstate.ClusterStateRegistry,
	nodeNameToNodeInfo map[string]*schedulernodeinfo.NodeInfo,
	nodes []*corev1.Node, sd *ScaleDown) (ScaleUpOptions, ScaleDownCandidates, error) {

	// construct requests
	req, err := GenerateAutoscalerRequest(context.CloudProvider.NodeGroups(), clusterStateRegistry.GetUpcomingNodes())
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot generate autoscaler requests, err: %s", err.Error())
	}

	// get configmap
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"Cannot marshal review to bytes, err: %s", err.Error())
	}
	reqData := string(b)
	cm, err := c.client.CoreV1().ConfigMaps(c.namespace).Get(c.name, metav1.GetOptions{})
	if err != nil && apierr.IsNotFound(err) {
		newcm := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: c.namespace,
				Name:      c.name,
			},
			Data: map[string]string{
				"request":  reqData,
				"response": "",
			},
		}
		_, createErr := c.client.CoreV1().ConfigMaps(c.namespace).Create(newcm)
		if createErr != nil {
			return nil, nil, fmt.Errorf(
				"Cannot create configmap %s/%s, err: %s", c.namespace, c.name, err.Error())
		}
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot get configmap %s/%s, err: %s",
			c.namespace, c.name, err.Error())
	}
	// update configmap
	cm.Data["request"] = reqData
	_, err = c.client.CoreV1().ConfigMaps(c.namespace).Update(cm)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot update configmap %s/%s, err: %s",
			c.namespace, c.name, err.Error())
	}

	// get response
	resData := cm.Data["response"]
	if len(resData) == 0 {
		klog.Infof("Response has not been written in configmap yet")
		return nil, nil, nil
	}
	var res AutoscalerResponse
	err = json.Unmarshal([]byte(resData), &res)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to unmarshal response to AutoscalerResponse, err: %s", err.Error())
	}
	faResp := ClusterAutoscalerReview{
		Request:  req,
		Response: &res,
	}
	klog.Infof("Get webhook response from configmap: %+v", res)
	options, candidates, err := HandleResponse(faResp, nodes, nodeNameToNodeInfo, sd)
	if err != nil {
		return nil, nil, err
	}

	return options, candidates, nil
}

// ExecuteScale execute scale up and down based on webhook responses
func (c *ConfigMapScaler) ExecuteScale(context *contextinternal.Context,
	clusterStateRegistry *clusterstate.ClusterStateRegistry,
	sd *ScaleDown, nodes []*corev1.Node, options ScaleUpOptions,
	candidates ScaleDownCandidates,
	nodeNameToNodeInfo map[string]*schedulernodeinfo.NodeInfo) error {

	err := ExecuteScaleUp(context, clusterStateRegistry, options)
	if err != nil {
		return err
	}

	if len(options) > 0 {
		klog.Infof("Scaling up node groups now, skip scaling down progress")
		return nil
	}

	err = ExecuteScaleDown(context, sd, nodes, candidates, nodeNameToNodeInfo)
	if err != nil {
		return err
	}

	return nil
}
