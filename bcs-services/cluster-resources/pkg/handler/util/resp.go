/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/structpb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	res "github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/resource"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/resource/formatter"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util"
)

func BuildListApiResp(
	clusterID, resKind, groupVersion, namespace string, opts metav1.ListOptions,
) (*structpb.Struct, error) {
	clusterConf := res.NewClusterConfig(clusterID)
	k8sRes, err := res.GenGroupVersionResource(clusterConf, clusterID, resKind, groupVersion)
	if err != nil {
		return nil, err
	}

	var ret *unstructured.UnstructuredList
	if namespace != "" {
		ret, err = res.ListNamespaceScopedRes(clusterConf, namespace, k8sRes, opts)
	} else {
		// TODO 支持集群域资源
		panic("cluster scoped resource unsupported")
	}
	if err != nil {
		return nil, err
	}

	manifest := ret.UnstructuredContent()
	manifestExt := map[string]interface{}{}
	// 遍历列表中的每个资源，生成 manifestExt
	for _, item := range manifest["items"].([]interface{}) {
		uid, _ := util.GetItems(item.(map[string]interface{}), "metadata.uid")
		manifestExt[uid.(string)] = formatter.Kind2FormatFuncMap[resKind](item.(map[string]interface{}))
	}
	respData := map[string]interface{}{"manifest": manifest, "manifestExt": manifestExt}
	return util.Map2pbStruct(respData)
}

func BuildRetrieveApiResp(
	clusterID, resKind, groupVersion, namespace, name string, opts metav1.GetOptions,
) (*structpb.Struct, error) {
	clusterConf := res.NewClusterConfig(clusterID)
	k8sRes, err := res.GenGroupVersionResource(clusterConf, clusterID, resKind, groupVersion)
	if err != nil {
		return nil, err
	}

	var ret *unstructured.Unstructured
	if namespace != "" {
		ret, err = res.GetNamespaceScopedRes(clusterConf, namespace, name, k8sRes, opts)
	} else {
		// TODO 支持集群域资源
		panic("cluster scoped resource unsupported")
	}
	if err != nil {
		return nil, err
	}

	manifest := ret.UnstructuredContent()
	respData := map[string]interface{}{
		"manifest": manifest, "manifestExt": formatter.Kind2FormatFuncMap[resKind](manifest),
	}
	return util.Map2pbStruct(respData)
}

func BuildCreateApiResp(
	clusterID, resKind, groupVersion string, manifest *structpb.Struct, isNamespaceScoped bool, opts metav1.CreateOptions,
) (*structpb.Struct, error) {
	clusterConf := res.NewClusterConfig(clusterID)
	k8sRes, err := res.GenGroupVersionResource(clusterConf, clusterID, resKind, groupVersion)
	if err != nil {
		return nil, err
	}

	var ret *unstructured.Unstructured
	if isNamespaceScoped {
		ret, err = res.CreateNamespaceScopedRes(clusterConf, manifest.AsMap(), k8sRes, opts)
	} else {
		// TODO 支持集群域资源
		panic("cluster scoped resource unsupported")
	}
	if err != nil {
		return nil, err
	}
	return util.Unstructured2pbStruct(ret), nil
}

func BuildUpdateApiResp(
	clusterID, resKind, groupVersion, namespace, name string, manifest *structpb.Struct, opts metav1.UpdateOptions,
) (*structpb.Struct, error) {
	clusterConf := res.NewClusterConfig(clusterID)
	k8sRes, err := res.GenGroupVersionResource(clusterConf, clusterID, resKind, groupVersion)
	if err != nil {
		return nil, err
	}

	var ret *unstructured.Unstructured
	if namespace != "" {
		ret, err = res.UpdateNamespaceScopedRes(clusterConf, namespace, name, manifest.AsMap(), k8sRes, opts)
	} else {
		// TODO 支持集群域资源
		panic("cluster scoped resource unsupported")
	}
	if err != nil {
		return nil, err
	}
	return util.Unstructured2pbStruct(ret), nil
}

func BuildDeleteApiResp(
	clusterID, resKind, groupVersion, namespace, name string, opts metav1.DeleteOptions,
) error {
	clusterConf := res.NewClusterConfig(clusterID)
	k8sRes, err := res.GenGroupVersionResource(clusterConf, clusterID, resKind, groupVersion)
	if err != nil {
		return err
	}
	if namespace != "" {
		return res.DeleteNamespaceScopedRes(clusterConf, namespace, name, k8sRes, opts)
	}
	// TODO 支持集群域资源
	panic("cluster scoped resource unsupported")
}

// 去除容器 ID 前缀，原格式：docker://[a-zA-Z0-9]{64}
func extractContainerID(rawContainerID string) string {
	return strings.Replace(rawContainerID, "docker://", "", 1)
}

func BuildListContainerApiResp(clusterID, namespace, podName string) (*structpb.ListValue, error) {
	podManifest, err := res.FetchPodManifest(clusterID, namespace, podName)
	if err != nil {
		return nil, err
	}

	containers := []map[string]interface{}{}
	containerStatuses, _ := util.GetItems(podManifest, "status.containerStatuses")
	for _, cs := range containerStatuses.([]interface{}) {
		cs, _ := cs.(map[string]interface{})
		status, reason, message := "", "", ""
		// state 有且只有一对键值：running / terminated / waiting
		// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#containerstate-v1-core
		for k := range cs["state"].(map[string]interface{}) {
			status = k
			reason, _ = util.GetWithDefault(cs, []string{"state", k, "reason"}, k).(string)
			message, _ = util.GetWithDefault(cs, []string{"state", k, "message"}, k).(string)
		}
		containers = append(containers, map[string]interface{}{
			"containerId": extractContainerID(util.GetWithDefault(cs, "containerID", "").(string)),
			"image":       cs["image"].(string),
			"name":        cs["name"].(string),
			"status":      status,
			"reason":      reason,
			"message":     message,
		})
	}
	return util.MapSlice2ListValue(containers)
}

func BuildGetContainerApiResp(clusterID, namespace, podName, containerName string) (*structpb.Struct, error) {
	podManifest, err := res.FetchPodManifest(clusterID, namespace, podName)
	if err != nil {
		return nil, err
	}

	// 遍历查找指定容器的 Spec 及 Status，若其中某项不存在，则抛出错误
	var curContainerStatus map[string]interface{}
	var curContainerSpec map[string]interface{}
	containerSpec, _ := util.GetItems(podManifest, "spec.containers")
	for _, csp := range containerSpec.([]interface{}) {
		csp, _ := csp.(map[string]interface{})
		if containerName == csp["name"].(string) {
			curContainerSpec = csp
		}
	}
	containerStatuses, _ := util.GetItems(podManifest, "status.containerStatuses")
	for _, cs := range containerStatuses.([]interface{}) {
		cs, _ := cs.(map[string]interface{})
		if containerName == cs["name"].(string) {
			curContainerStatus = cs
		}
	}
	if len(curContainerSpec) == 0 || len(curContainerStatus) == 0 {
		return nil, fmt.Errorf("container %s spec or status not found", containerName)
	}

	containerInfo := map[string]interface{}{
		"hostName":      util.GetWithDefault(podManifest, "spec.nodeName", "--"),
		"hostIP":        util.GetWithDefault(podManifest, "status.hostIP", "--"),
		"containerIP":   util.GetWithDefault(podManifest, "status.podIP", "--"),
		"containerID":   extractContainerID(util.GetWithDefault(curContainerStatus, "containerID", "").(string)),
		"containerName": containerName,
		"image":         util.GetWithDefault(curContainerStatus, "image", "--"),
		"networkMode":   util.GetWithDefault(podManifest, "spec.dnsPolicy", "--"),
		"ports":         util.GetWithDefault(curContainerSpec, "ports", []interface{}{}),
		"volumes":       []map[string]interface{}{},
		"resources":     util.GetWithDefault(curContainerSpec, "resources", map[string]interface{}{}),
		"command": map[string]interface{}{
			"command": util.GetWithDefault(curContainerSpec, "command", []string{}),
			"args":    util.GetWithDefault(curContainerSpec, "args", []string{}),
		},
	}
	mounts := util.GetWithDefault(curContainerSpec, "volumeMounts", []map[string]interface{}{})
	for _, m := range mounts.([]interface{}) {
		m, _ := m.(map[string]interface{})
		containerInfo["volumes"] = append(containerInfo["volumes"].([]map[string]interface{}), map[string]interface{}{
			"name":      util.GetWithDefault(m, "name", "--"),
			"mountPath": util.GetWithDefault(m, "mountPath", "--"),
			"readonly":  util.GetWithDefault(m, "readOnly", "--"),
		})
	}

	return util.Map2pbStruct(containerInfo)
}
