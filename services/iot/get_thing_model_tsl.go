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

// GetThingModelTsl invokes the iot.GetThingModelTsl API synchronously
func (client *Client) GetThingModelTsl(request *GetThingModelTslRequest) (response *GetThingModelTslResponse, err error) {
	response = CreateGetThingModelTslResponse()
	err = client.DoAction(request, response)
	return
}

// GetThingModelTslWithChan invokes the iot.GetThingModelTsl API asynchronously
func (client *Client) GetThingModelTslWithChan(request *GetThingModelTslRequest) (<-chan *GetThingModelTslResponse, <-chan error) {
	responseChan := make(chan *GetThingModelTslResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetThingModelTsl(request)
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

// GetThingModelTslWithCallback invokes the iot.GetThingModelTsl API asynchronously
func (client *Client) GetThingModelTslWithCallback(request *GetThingModelTslRequest, callback func(response *GetThingModelTslResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetThingModelTslResponse
		var err error
		defer close(result)
		response, err = client.GetThingModelTsl(request)
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

// GetThingModelTslRequest is the request struct for api GetThingModelTsl
type GetThingModelTslRequest struct {
	*requests.RpcRequest
	RealTenantId      string           `position:"Query" name:"RealTenantId"`
	Simple            requests.Boolean `position:"Query" name:"Simple"`
	RealTripartiteKey string           `position:"Query" name:"RealTripartiteKey"`
	DTModelId         string           `position:"Query" name:"DTModelId"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	ModelVersion      string           `position:"Query" name:"ModelVersion"`
	FunctionBlockId   string           `position:"Query" name:"FunctionBlockId"`
}

// GetThingModelTslResponse is the response struct for api GetThingModelTsl
type GetThingModelTslResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGetThingModelTslRequest creates a request to invoke GetThingModelTsl API
func CreateGetThingModelTslRequest() (request *GetThingModelTslRequest) {
	request = &GetThingModelTslRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetThingModelTsl", "", "")
	request.Method = requests.POST
	return
}

// CreateGetThingModelTslResponse creates a response to parse from GetThingModelTsl response
func CreateGetThingModelTslResponse() (response *GetThingModelTslResponse) {
	response = &GetThingModelTslResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
