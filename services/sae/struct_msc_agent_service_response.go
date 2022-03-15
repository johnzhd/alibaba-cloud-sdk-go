package sae

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// MscAgentServiceResponse is a nested struct in sae response
type MscAgentServiceResponse struct {
	EdasAppName string `json:"EdasAppName" xml:"EdasAppName"`
	Version     string `json:"Version" xml:"Version"`
	InstanceNum int64  `json:"InstanceNum" xml:"InstanceNum"`
	EdasAppId   string `json:"EdasAppId" xml:"EdasAppId"`
	ServiceName string `json:"ServiceName" xml:"ServiceName"`
	Group       string `json:"Group" xml:"Group"`
}
