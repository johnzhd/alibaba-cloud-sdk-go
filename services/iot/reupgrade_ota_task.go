package iot

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

// ReupgradeOTATask invokes the iot.ReupgradeOTATask API synchronously
func (client *Client) ReupgradeOTATask(request *ReupgradeOTATaskRequest) (response *ReupgradeOTATaskResponse, err error) {
	response = CreateReupgradeOTATaskResponse()
	err = client.DoAction(request, response)
	return
}

// ReupgradeOTATaskWithChan invokes the iot.ReupgradeOTATask API asynchronously
func (client *Client) ReupgradeOTATaskWithChan(request *ReupgradeOTATaskRequest) (<-chan *ReupgradeOTATaskResponse, <-chan error) {
	responseChan := make(chan *ReupgradeOTATaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReupgradeOTATask(request)
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

// ReupgradeOTATaskWithCallback invokes the iot.ReupgradeOTATask API asynchronously
func (client *Client) ReupgradeOTATaskWithCallback(request *ReupgradeOTATaskRequest, callback func(response *ReupgradeOTATaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReupgradeOTATaskResponse
		var err error
		defer close(result)
		response, err = client.ReupgradeOTATask(request)
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

// ReupgradeOTATaskRequest is the request struct for api ReupgradeOTATask
type ReupgradeOTATaskRequest struct {
	*requests.RpcRequest
	JobId         string    `position:"Query" name:"JobId"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	TaskId        *[]string `position:"Query" name:"TaskId"  type:"Repeated"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// ReupgradeOTATaskResponse is the response struct for api ReupgradeOTATask
type ReupgradeOTATaskResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateReupgradeOTATaskRequest creates a request to invoke ReupgradeOTATask API
func CreateReupgradeOTATaskRequest() (request *ReupgradeOTATaskRequest) {
	request = &ReupgradeOTATaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ReupgradeOTATask", "", "")
	request.Method = requests.POST
	return
}

// CreateReupgradeOTATaskResponse creates a response to parse from ReupgradeOTATask response
func CreateReupgradeOTATaskResponse() (response *ReupgradeOTATaskResponse) {
	response = &ReupgradeOTATaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
