/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package list

import (
	"context"
	"encoding/json"
	"fmt"

	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/app/api/proto/logmanager"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/logmanager/v1"
)

func listLogCollectionTask(c *utils.ClientContext) error {
	manager, err := v1.NewLogManager(context.Background(), utils.GetClientOption())
	if err != nil {
		return err
	}
	var req proto.ListLogCollectionTaskReq
	if c.IsSet(utils.OptionFile) {
		var data []byte
		data, err = c.FileData()
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &req)
		if err != nil {
			return err
		}
	} else {
		req.ClusterIDs = c.ClusterID()
		if !c.IsAllNamespace() {
			req.ConfigNamespace = c.Namespace()
		}
	}
	configs, err := manager.ListLogCollectionTask(&req)
	if err != nil {
		return err
	}
	data, err := json.Marshal(configs)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(data))
	return nil
}
