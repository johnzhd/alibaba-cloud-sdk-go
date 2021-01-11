package dts

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

// SynchronizationInstance is a nested struct in dts response
type SynchronizationInstance struct {
	CreateTime                    string                                      `json:"CreateTime" xml:"CreateTime"`
	DataInitialization            string                                      `json:"DataInitialization" xml:"DataInitialization"`
	Delay                         string                                      `json:"Delay" xml:"Delay"`
	ErrorMessage                  string                                      `json:"ErrorMessage" xml:"ErrorMessage"`
	ExpireTime                    string                                      `json:"ExpireTime" xml:"ExpireTime"`
	PayType                       string                                      `json:"PayType" xml:"PayType"`
	Status                        string                                      `json:"Status" xml:"Status"`
	StructureInitialization       string                                      `json:"StructureInitialization" xml:"StructureInitialization"`
	SynchronizationDirection      string                                      `json:"SynchronizationDirection" xml:"SynchronizationDirection"`
	SynchronizationJobClass       string                                      `json:"SynchronizationJobClass" xml:"SynchronizationJobClass"`
	SynchronizationJobId          string                                      `json:"SynchronizationJobId" xml:"SynchronizationJobId"`
	SynchronizationJobName        string                                      `json:"SynchronizationJobName" xml:"SynchronizationJobName"`
	DataInitializationStatus      DataInitializationStatus                    `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	DataSynchronizationStatus     DataSynchronizationStatus                   `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	DestinationEndpoint           DestinationEndpoint                         `json:"DestinationEndpoint" xml:"DestinationEndpoint"`
	Performance                   Performance                                 `json:"Performance" xml:"Performance"`
	PrecheckStatus                PrecheckStatusInDescribeSynchronizationJobs `json:"PrecheckStatus" xml:"PrecheckStatus"`
	SourceEndpoint                SourceEndpoint                              `json:"SourceEndpoint" xml:"SourceEndpoint"`
	StructureInitializationStatus StructureInitializationStatus               `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
	Tags                          []Tag                                       `json:"Tags" xml:"Tags"`
	SynchronizationObjects        []SynchronizationObject                     `json:"SynchronizationObjects" xml:"SynchronizationObjects"`
}
