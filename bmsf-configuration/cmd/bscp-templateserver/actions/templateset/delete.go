/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package templateset

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"

	pbcommon "bk-bscp/internal/protocol/common"
	pbdatamanager "bk-bscp/internal/protocol/datamanager"
	pb "bk-bscp/internal/protocol/templateserver"
	"bk-bscp/pkg/common"
	"bk-bscp/pkg/logger"
)

// DeleteAction delete a template set
type DeleteAction struct {
	viper      *viper.Viper
	dataMgrCli pbdatamanager.DataManagerClient

	req  *pb.DeleteConfigTemplateSetReq
	resp *pb.DeleteConfigTemplateSetResp
}

// NewDeleteAction create new DeleteAction
func NewDeleteAction(viper *viper.Viper, dataMgrCli pbdatamanager.DataManagerClient,
	req *pb.DeleteConfigTemplateSetReq, resp *pb.DeleteConfigTemplateSetResp) *DeleteAction {
	action := &DeleteAction{viper: viper, dataMgrCli: dataMgrCli, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.ErrCode = pbcommon.ErrCode_E_OK
	action.resp.ErrMsg = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *DeleteAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.ErrCode = errCode
	act.resp.ErrMsg = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *DeleteAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_BS_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *DeleteAction) Output() error {
	// do nothing.
	return nil
}

func (act *DeleteAction) verify() error {
	if err := common.VerifyID(act.req.Bid, "bid"); err != nil {
		return nil
	}

	if err := common.VerifyID(act.req.Setid, "setid"); err != nil {
		return nil
	}

	if err := common.VerifyNormalName(act.req.Operator, "operator"); err != nil {
		return nil
	}

	return nil
}

func (act *DeleteAction) deleteTemplateSet() (pbcommon.ErrCode, string) {
	req := &pbdatamanager.DeleteConfigTemplateSetReq{
		Seq:      act.req.Seq,
		Bid:      act.req.Bid,
		Setid:    act.req.Setid,
		Operator: act.req.Operator,
	}

	ctx, cancel := context.WithTimeout(context.Background(), act.viper.GetDuration("datamanager.calltimeout"))
	defer cancel()

	logger.V(2).Infof("DeleteConfigTemplateSet[%d]|request to datamanager DeleteConfigTemplateSet, %+v", act.req.Seq, req)

	resp, err := act.dataMgrCli.DeleteConfigTemplateSet(ctx, req)
	if err != nil {
		return pbcommon.ErrCode_E_TPL_SYSTEM_UNKONW, fmt.Sprintf("request to datamanager DeleteConfigTemplateSet, %+v", err)
	}

	if resp.ErrCode != pbcommon.ErrCode_E_OK {
		return resp.ErrCode, resp.ErrMsg
	}

	return pbcommon.ErrCode_E_OK, "OK"
}

// Do do action
func (act *DeleteAction) Do() error {
	// delete config template set.
	if errCode, errMsg := act.deleteTemplateSet(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
