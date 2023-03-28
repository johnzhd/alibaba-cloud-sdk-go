package alinlp

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

// GetOperationChMedical invokes the alinlp.GetOperationChMedical API synchronously
func (client *Client) GetOperationChMedical(request *GetOperationChMedicalRequest) (response *GetOperationChMedicalResponse, err error) {
	response = CreateGetOperationChMedicalResponse()
	err = client.DoAction(request, response)
	return
}

// GetOperationChMedicalWithChan invokes the alinlp.GetOperationChMedical API asynchronously
func (client *Client) GetOperationChMedicalWithChan(request *GetOperationChMedicalRequest) (<-chan *GetOperationChMedicalResponse, <-chan error) {
	responseChan := make(chan *GetOperationChMedicalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOperationChMedical(request)
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

// GetOperationChMedicalWithCallback invokes the alinlp.GetOperationChMedical API asynchronously
func (client *Client) GetOperationChMedicalWithCallback(request *GetOperationChMedicalRequest, callback func(response *GetOperationChMedicalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOperationChMedicalResponse
		var err error
		defer close(result)
		response, err = client.GetOperationChMedical(request)
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

// GetOperationChMedicalRequest is the request struct for api GetOperationChMedical
type GetOperationChMedicalRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Name        string `position:"Body" name:"Name"`
}

// GetOperationChMedicalResponse is the response struct for api GetOperationChMedical
type GetOperationChMedicalResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetOperationChMedicalRequest creates a request to invoke GetOperationChMedical API
func CreateGetOperationChMedicalRequest() (request *GetOperationChMedicalRequest) {
	request = &GetOperationChMedicalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetOperationChMedical", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOperationChMedicalResponse creates a response to parse from GetOperationChMedical response
func CreateGetOperationChMedicalResponse() (response *GetOperationChMedicalResponse) {
	response = &GetOperationChMedicalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
