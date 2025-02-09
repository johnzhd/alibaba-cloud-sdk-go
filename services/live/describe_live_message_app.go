package live

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

// DescribeLiveMessageApp invokes the live.DescribeLiveMessageApp API synchronously
func (client *Client) DescribeLiveMessageApp(request *DescribeLiveMessageAppRequest) (response *DescribeLiveMessageAppResponse, err error) {
	response = CreateDescribeLiveMessageAppResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveMessageAppWithChan invokes the live.DescribeLiveMessageApp API asynchronously
func (client *Client) DescribeLiveMessageAppWithChan(request *DescribeLiveMessageAppRequest) (<-chan *DescribeLiveMessageAppResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveMessageAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveMessageApp(request)
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

// DescribeLiveMessageAppWithCallback invokes the live.DescribeLiveMessageApp API asynchronously
func (client *Client) DescribeLiveMessageAppWithCallback(request *DescribeLiveMessageAppRequest, callback func(response *DescribeLiveMessageAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveMessageAppResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveMessageApp(request)
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

// DescribeLiveMessageAppRequest is the request struct for api DescribeLiveMessageApp
type DescribeLiveMessageAppRequest struct {
	*requests.RpcRequest
	DataCenter string `position:"Query" name:"DataCenter"`
	AppId      string `position:"Query" name:"AppId"`
}

// DescribeLiveMessageAppResponse is the response struct for api DescribeLiveMessageApp
type DescribeLiveMessageAppResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	AppId       string `json:"AppId" xml:"AppId"`
	AppKey      string `json:"AppKey" xml:"AppKey"`
	AppSign     string `json:"AppSign" xml:"AppSign"`
	AuditType   int    `json:"AuditType" xml:"AuditType"`
	AuditUrl    string `json:"AuditUrl" xml:"AuditUrl"`
	CallbackUrl string `json:"CallbackUrl" xml:"CallbackUrl"`
	Disable     bool   `json:"Disable" xml:"Disable"`
}

// CreateDescribeLiveMessageAppRequest creates a request to invoke DescribeLiveMessageApp API
func CreateDescribeLiveMessageAppRequest() (request *DescribeLiveMessageAppRequest) {
	request = &DescribeLiveMessageAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveMessageApp", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveMessageAppResponse creates a response to parse from DescribeLiveMessageApp response
func CreateDescribeLiveMessageAppResponse() (response *DescribeLiveMessageAppResponse) {
	response = &DescribeLiveMessageAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
