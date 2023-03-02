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

// QueryProductTopic invokes the iot.QueryProductTopic API synchronously
func (client *Client) QueryProductTopic(request *QueryProductTopicRequest) (response *QueryProductTopicResponse, err error) {
	response = CreateQueryProductTopicResponse()
	err = client.DoAction(request, response)
	return
}

// QueryProductTopicWithChan invokes the iot.QueryProductTopic API asynchronously
func (client *Client) QueryProductTopicWithChan(request *QueryProductTopicRequest) (<-chan *QueryProductTopicResponse, <-chan error) {
	responseChan := make(chan *QueryProductTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryProductTopic(request)
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

// QueryProductTopicWithCallback invokes the iot.QueryProductTopic API asynchronously
func (client *Client) QueryProductTopicWithCallback(request *QueryProductTopicRequest, callback func(response *QueryProductTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryProductTopicResponse
		var err error
		defer close(result)
		response, err = client.QueryProductTopic(request)
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

// QueryProductTopicRequest is the request struct for api QueryProductTopic
type QueryProductTopicRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// QueryProductTopicResponse is the response struct for api QueryProductTopic
type QueryProductTopicResponse struct {
	*responses.BaseResponse
	RequestId    string                  `json:"RequestId" xml:"RequestId"`
	Success      bool                    `json:"Success" xml:"Success"`
	Code         string                  `json:"Code" xml:"Code"`
	ErrorMessage string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryProductTopic `json:"Data" xml:"Data"`
}

// CreateQueryProductTopicRequest creates a request to invoke QueryProductTopic API
func CreateQueryProductTopicRequest() (request *QueryProductTopicRequest) {
	request = &QueryProductTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryProductTopic", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryProductTopicResponse creates a response to parse from QueryProductTopic response
func CreateQueryProductTopicResponse() (response *QueryProductTopicResponse) {
	response = &QueryProductTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
