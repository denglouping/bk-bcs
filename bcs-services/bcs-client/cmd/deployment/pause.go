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

package deployment

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	v4 "github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"
)

// NewPauseCommand create deployment pause command
func NewPauseCommand() cli.Command {
	return cli.Command{
		Name:  "pause",
		Usage: "pause deployment update",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Usage: "Pause type, deployment",
			},
			cli.StringFlag{
				Name:  "clusterid",
				Usage: "Cluster ID",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Deployment name",
			},
			cli.StringFlag{
				Name:  "namespace, ns",
				Usage: "Namespace",
				Value: "defaultGroup",
			},
		},
		Action: func(c *cli.Context) error {
			if err := pause(utils.NewClientContext(c)); err != nil {
				return err
			}
			return nil
		},
	}
}

func pause(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionType); err != nil {
		return err
	}

	resourceType := c.String(utils.OptionType)

	switch resourceType {
	case deploy, deployment:
		return pauseDeployment(c)
	default:
		return fmt.Errorf("invalid type: %s", resourceType)
	}
}

func pauseDeployment(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID, utils.OptionNamespace, utils.OptionName); err != nil {
		return err
	}

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())
	err := scheduler.PauseDeployment(c.ClusterID(), c.Namespace(), c.String(utils.OptionName))
	if err != nil {
		return fmt.Errorf("failed to pause deployment: %v", err)
	}

	fmt.Printf("success to pause deployment\n")
	return nil
}
