/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package yunikorn

import (
	"fmt"

	"github.com/apache/yunikorn-k8shim/test/e2e/framework/configmanager"
)

type ResourceUsage struct {
	memory int64
	vCPU   int64
}

func (r *ResourceUsage) ParseResourceUsage(resource map[string]interface{}) {
	r.memory = int64(resource["memory"].(float64))
	r.vCPU = int64(resource["vcore"].(float64))
}

func (r *ResourceUsage) GetMemory() int64 {
	return r.memory
}

func (r *ResourceUsage) GetVCPU() int64 {
	return r.vCPU
}

func GetYKUrl() string {
	return fmt.Sprintf("%s://%s",
		configmanager.YuniKornTestConfig.YkScheme,
		GetYKHost(),
	)
}

func GetYKHost() string {
	return fmt.Sprintf("%s:%s",
		configmanager.YuniKornTestConfig.YkHost,
		configmanager.YuniKornTestConfig.YkPort,
	)
}

func GetYKScheme() string {
	return configmanager.YuniKornTestConfig.YkScheme
}
