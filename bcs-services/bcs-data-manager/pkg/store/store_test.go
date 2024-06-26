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

package store

import (
	"testing"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/odm/drivers/mongo"
	"github.com/stretchr/testify/assert"
)

func TestServer_initModel(t *testing.T) {

	mongoOptions := &mongo.Options{
		Hosts:                 []string{"127.0.0.1:27017"},
		ConnectTimeoutSeconds: 3,
		Database:              "datamanager_test",
		Username:              "data",
		Password:              "test1234",
	}
	mongoDB, err := mongo.NewDB(mongoOptions)
	assert.Equal(t, nil, err)
	err = mongoDB.Ping()
	assert.Equal(t, nil, err)
}
