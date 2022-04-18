package sddp

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

// DescribeEventTypes invokes the sddp.DescribeEventTypes API synchronously
func (client *Client) DescribeEventTypes(request *DescribeEventTypesRequest) (response *DescribeEventTypesResponse, err error) {
	response = CreateDescribeEventTypesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEventTypesWithChan invokes the sddp.DescribeEventTypes API asynchronously
func (client *Client) DescribeEventTypesWithChan(request *DescribeEventTypesRequest) (<-chan *DescribeEventTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeEventTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEventTypes(request)
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

// DescribeEventTypesWithCallback invokes the sddp.DescribeEventTypes API asynchronously
func (client *Client) DescribeEventTypesWithCallback(request *DescribeEventTypesRequest, callback func(response *DescribeEventTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEventTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeEventTypes(request)
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

// DescribeEventTypesRequest is the request struct for api DescribeEventTypes
type DescribeEventTypesRequest struct {
	*requests.RpcRequest
	ResourceId   requests.Integer `position:"Query" name:"ResourceId"`
	ParentTypeId requests.Integer `position:"Query" name:"ParentTypeId"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	Lang         string           `position:"Query" name:"Lang"`
	Status       requests.Integer `position:"Query" name:"Status"`
}

// DescribeEventTypesResponse is the response struct for api DescribeEventTypes
type DescribeEventTypesResponse struct {
	*responses.BaseResponse
	RequestId     string      `json:"RequestId" xml:"RequestId"`
	EventTypeList []EventType `json:"EventTypeList" xml:"EventTypeList"`
}

// CreateDescribeEventTypesRequest creates a request to invoke DescribeEventTypes API
func CreateDescribeEventTypesRequest() (request *DescribeEventTypesRequest) {
	request = &DescribeEventTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeEventTypes", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEventTypesResponse creates a response to parse from DescribeEventTypes response
func CreateDescribeEventTypesResponse() (response *DescribeEventTypesResponse) {
	response = &DescribeEventTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
