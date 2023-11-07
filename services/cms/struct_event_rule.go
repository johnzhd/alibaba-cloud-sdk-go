package cms

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

// EventRule is a nested struct in cms response
type EventRule struct {
	EventType    string                              `json:"EventType" xml:"EventType"`
	Description  string                              `json:"Description" xml:"Description"`
	GroupId      string                              `json:"GroupId" xml:"GroupId"`
	Name         string                              `json:"Name" xml:"Name"`
	State        string                              `json:"State" xml:"State"`
	SilenceTime  int64                               `json:"SilenceTime" xml:"SilenceTime"`
	EventPattern EventPatternItem `json:"EventPattern" xml:"EventPattern"`
}
