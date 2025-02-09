package cloud_siem

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

// DoSelfDelegate invokes the cloud_siem.DoSelfDelegate API synchronously
func (client *Client) DoSelfDelegate(request *DoSelfDelegateRequest) (response *DoSelfDelegateResponse, err error) {
	response = CreateDoSelfDelegateResponse()
	err = client.DoAction(request, response)
	return
}

// DoSelfDelegateWithChan invokes the cloud_siem.DoSelfDelegate API asynchronously
func (client *Client) DoSelfDelegateWithChan(request *DoSelfDelegateRequest) (<-chan *DoSelfDelegateResponse, <-chan error) {
	responseChan := make(chan *DoSelfDelegateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DoSelfDelegate(request)
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

// DoSelfDelegateWithCallback invokes the cloud_siem.DoSelfDelegate API asynchronously
func (client *Client) DoSelfDelegateWithCallback(request *DoSelfDelegateRequest, callback func(response *DoSelfDelegateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DoSelfDelegateResponse
		var err error
		defer close(result)
		response, err = client.DoSelfDelegate(request)
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

// DoSelfDelegateRequest is the request struct for api DoSelfDelegate
type DoSelfDelegateRequest struct {
	*requests.RpcRequest
	AliUid        requests.Integer `position:"Body" name:"AliUid"`
	DelegateOrNot requests.Integer `position:"Body" name:"DelegateOrNot"`
}

// DoSelfDelegateResponse is the response struct for api DoSelfDelegate
type DoSelfDelegateResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	DyCode    string `json:"DyCode" xml:"DyCode"`
	DyMessage string `json:"DyMessage" xml:"DyMessage"`
}

// CreateDoSelfDelegateRequest creates a request to invoke DoSelfDelegate API
func CreateDoSelfDelegateRequest() (request *DoSelfDelegateRequest) {
	request = &DoSelfDelegateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DoSelfDelegate", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDoSelfDelegateResponse creates a response to parse from DoSelfDelegate response
func CreateDoSelfDelegateResponse() (response *DoSelfDelegateResponse) {
	response = &DoSelfDelegateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
