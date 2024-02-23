//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func Test_Mock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	cs := cloudstack.NewMockClient(mockCtrl)
	ms := cs.Host.(*cloudstack.MockHostServiceIface)

	ms.EXPECT().GetHostByID("host-id").Times(1).Return(
		&cloudstack.Host{
			Id: "host-id",
		}, 2, nil)

	cs.Host.GetHostByID("host-id")
}
