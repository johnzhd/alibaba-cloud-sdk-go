package dypnsapi

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

// GetAuthToken invokes the dypnsapi.GetAuthToken API synchronously
func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (response *GetAuthTokenResponse, err error) {
	response = CreateGetAuthTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuthTokenWithChan invokes the dypnsapi.GetAuthToken API asynchronously
func (client *Client) GetAuthTokenWithChan(request *GetAuthTokenRequest) (<-chan *GetAuthTokenResponse, <-chan error) {
	responseChan := make(chan *GetAuthTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuthToken(request)
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

// GetAuthTokenWithCallback invokes the dypnsapi.GetAuthToken API asynchronously
func (client *Client) GetAuthTokenWithCallback(request *GetAuthTokenRequest, callback func(response *GetAuthTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuthTokenResponse
		var err error
		defer close(result)
		response, err = client.GetAuthToken(request)
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

// GetAuthTokenRequest is the request struct for api GetAuthToken
type GetAuthTokenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Origin               string           `position:"Query" name:"Origin"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Url                  string           `position:"Query" name:"Url"`
}

// GetAuthTokenResponse is the response struct for api GetAuthToken
type GetAuthTokenResponse struct {
	*responses.BaseResponse
	Code      string    `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	RequestId string    `json:"RequestId" xml:"RequestId"`
	TokenInfo TokenInfo `json:"TokenInfo" xml:"TokenInfo"`
}

// CreateGetAuthTokenRequest creates a request to invoke GetAuthToken API
func CreateGetAuthTokenRequest() (request *GetAuthTokenRequest) {
	request = &GetAuthTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "GetAuthToken", "dypnsapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAuthTokenResponse creates a response to parse from GetAuthToken response
func CreateGetAuthTokenResponse() (response *GetAuthTokenResponse) {
	response = &GetAuthTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
