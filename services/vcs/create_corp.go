package vcs

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

// CreateCorp invokes the vcs.CreateCorp API synchronously
// api document: https://help.aliyun.com/api/vcs/createcorp.html
func (client *Client) CreateCorp(request *CreateCorpRequest) (response *CreateCorpResponse, err error) {
	response = CreateCreateCorpResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCorpWithChan invokes the vcs.CreateCorp API asynchronously
// api document: https://help.aliyun.com/api/vcs/createcorp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCorpWithChan(request *CreateCorpRequest) (<-chan *CreateCorpResponse, <-chan error) {
	responseChan := make(chan *CreateCorpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCorp(request)
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

// CreateCorpWithCallback invokes the vcs.CreateCorp API asynchronously
// api document: https://help.aliyun.com/api/vcs/createcorp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCorpWithCallback(request *CreateCorpRequest, callback func(response *CreateCorpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCorpResponse
		var err error
		defer close(result)
		response, err = client.CreateCorp(request)
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

// CreateCorpRequest is the request struct for api CreateCorp
type CreateCorpRequest struct {
	*requests.RpcRequest
	ParentCorpId string `position:"Body" name:"ParentCorpId"`
	Description  string `position:"Body" name:"Description"`
	AppName      string `position:"Body" name:"AppName"`
	CorpName     string `position:"Body" name:"CorpName"`
}

// CreateCorpResponse is the response struct for api CreateCorp
type CreateCorpResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	CorpId    string `json:"CorpId" xml:"CorpId"`
}

// CreateCreateCorpRequest creates a request to invoke CreateCorp API
func CreateCreateCorpRequest() (request *CreateCorpRequest) {
	request = &CreateCorpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "CreateCorp", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCorpResponse creates a response to parse from CreateCorp response
func CreateCreateCorpResponse() (response *CreateCorpResponse) {
	response = &CreateCorpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
