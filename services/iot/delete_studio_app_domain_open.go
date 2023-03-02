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

// DeleteStudioAppDomainOpen invokes the iot.DeleteStudioAppDomainOpen API synchronously
func (client *Client) DeleteStudioAppDomainOpen(request *DeleteStudioAppDomainOpenRequest) (response *DeleteStudioAppDomainOpenResponse, err error) {
	response = CreateDeleteStudioAppDomainOpenResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteStudioAppDomainOpenWithChan invokes the iot.DeleteStudioAppDomainOpen API asynchronously
func (client *Client) DeleteStudioAppDomainOpenWithChan(request *DeleteStudioAppDomainOpenRequest) (<-chan *DeleteStudioAppDomainOpenResponse, <-chan error) {
	responseChan := make(chan *DeleteStudioAppDomainOpenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteStudioAppDomainOpen(request)
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

// DeleteStudioAppDomainOpenWithCallback invokes the iot.DeleteStudioAppDomainOpen API asynchronously
func (client *Client) DeleteStudioAppDomainOpenWithCallback(request *DeleteStudioAppDomainOpenRequest, callback func(response *DeleteStudioAppDomainOpenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteStudioAppDomainOpenResponse
		var err error
		defer close(result)
		response, err = client.DeleteStudioAppDomainOpen(request)
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

// DeleteStudioAppDomainOpenRequest is the request struct for api DeleteStudioAppDomainOpen
type DeleteStudioAppDomainOpenRequest struct {
	*requests.RpcRequest
	DomainId      requests.Integer `position:"Body" name:"DomainId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	ProjectId     string           `position:"Body" name:"ProjectId"`
	AppId         string           `position:"Body" name:"AppId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// DeleteStudioAppDomainOpenResponse is the response struct for api DeleteStudioAppDomainOpen
type DeleteStudioAppDomainOpenResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         bool   `json:"Data" xml:"Data"`
}

// CreateDeleteStudioAppDomainOpenRequest creates a request to invoke DeleteStudioAppDomainOpen API
func CreateDeleteStudioAppDomainOpenRequest() (request *DeleteStudioAppDomainOpenRequest) {
	request = &DeleteStudioAppDomainOpenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteStudioAppDomainOpen", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteStudioAppDomainOpenResponse creates a response to parse from DeleteStudioAppDomainOpen response
func CreateDeleteStudioAppDomainOpenResponse() (response *DeleteStudioAppDomainOpenResponse) {
	response = &DeleteStudioAppDomainOpenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
