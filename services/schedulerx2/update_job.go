package schedulerx2

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateJob invokes the schedulerx2.UpdateJob API synchronously
func (client *Client) UpdateJob(request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
	response = CreateUpdateJobResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateJobWithChan invokes the schedulerx2.UpdateJob API asynchronously
func (client *Client) UpdateJobWithChan(request *UpdateJobRequest) (<-chan *UpdateJobResponse, <-chan error) {
	responseChan := make(chan *UpdateJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateJob(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateJobWithCallback invokes the schedulerx2.UpdateJob API asynchronously
func (client *Client) UpdateJobWithCallback(request *UpdateJobRequest, callback func(response *UpdateJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateJobResponse
		var err error
		defer close(result)
		response, err = client.UpdateJob(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateJobRequest is the request struct for api UpdateJob
type UpdateJobRequest struct {
	*requests.RpcRequest
	AttemptInterval     requests.Integer        `position:"Body" name:"AttemptInterval"`
	FailTimes           requests.Integer        `position:"Body" name:"FailTimes"`
	JobId               requests.Integer        `position:"Body" name:"JobId"`
	ConsumerSize        requests.Integer        `position:"Body" name:"ConsumerSize"`
	JarUrl              string                  `position:"Body" name:"JarUrl"`
	GroupId             string                  `position:"Body" name:"GroupId"`
	TaskMaxAttempt      requests.Integer        `position:"Body" name:"TaskMaxAttempt"`
	DataOffset          requests.Integer        `position:"Body" name:"DataOffset"`
	DispatcherSize      requests.Integer        `position:"Body" name:"DispatcherSize"`
	TaskAttemptInterval requests.Integer        `position:"Body" name:"TaskAttemptInterval"`
	ExecuteMode         string                  `position:"Body" name:"ExecuteMode"`
	TimeExpression      string                  `position:"Body" name:"TimeExpression"`
	TimeoutEnable       requests.Boolean        `position:"Body" name:"TimeoutEnable"`
	ContactInfo         *[]UpdateJobContactInfo `position:"Body" name:"ContactInfo"  type:"Repeated"`
	Name                string                  `position:"Body" name:"Name"`
	TimeType            requests.Integer        `position:"Body" name:"TimeType"`
	Parameters          string                  `position:"Body" name:"Parameters"`
	NamespaceSource     string                  `position:"Body" name:"NamespaceSource"`
	Timezone            string                  `position:"Body" name:"Timezone"`
	Description         string                  `position:"Body" name:"Description"`
	Content             string                  `position:"Body" name:"Content"`
	Timeout             requests.Integer        `position:"Body" name:"Timeout"`
	TimeoutKillEnable   requests.Boolean        `position:"Body" name:"TimeoutKillEnable"`
	PageSize            requests.Integer        `position:"Body" name:"PageSize"`
	TaskDispatchMode    string                  `position:"Body" name:"TaskDispatchMode"`
	Calendar            string                  `position:"Body" name:"Calendar"`
	FailEnable          requests.Boolean        `position:"Body" name:"FailEnable"`
	SendChannel         string                  `position:"Body" name:"SendChannel"`
	MaxAttempt          requests.Integer        `position:"Body" name:"MaxAttempt"`
	MissWorkerEnable    requests.Boolean        `position:"Body" name:"MissWorkerEnable"`
	SuccessNoticeEnable requests.Boolean        `position:"Body" name:"SuccessNoticeEnable"`
	QueueSize           requests.Integer        `position:"Body" name:"QueueSize"`
	ClassName           string                  `position:"Body" name:"ClassName"`
	Namespace           string                  `position:"Body" name:"Namespace"`
	MaxConcurrency      requests.Integer        `position:"Body" name:"MaxConcurrency"`
}

// UpdateJobContactInfo is a repeated param struct in UpdateJobRequest
type UpdateJobContactInfo struct {
	Ding      string `name:"Ding"`
	UserPhone string `name:"UserPhone"`
	UserMail  string `name:"UserMail"`
	UserName  string `name:"UserName"`
}

// UpdateJobResponse is the response struct for api UpdateJob
type UpdateJobResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateJobRequest creates a request to invoke UpdateJob API
func CreateUpdateJobRequest() (request *UpdateJobRequest) {
	request = &UpdateJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "UpdateJob", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateJobResponse creates a response to parse from UpdateJob response
func CreateUpdateJobResponse() (response *UpdateJobResponse) {
	response = &UpdateJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
