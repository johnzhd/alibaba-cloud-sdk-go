package amp

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

// Progress invokes the amp.Progress API synchronously
func (client *Client) Progress(request *ProgressRequest) (response *ProgressResponse, err error) {
	response = CreateProgressResponse()
	err = client.DoAction(request, response)
	return
}

// ProgressWithChan invokes the amp.Progress API asynchronously
func (client *Client) ProgressWithChan(request *ProgressRequest) (<-chan *ProgressResponse, <-chan error) {
	responseChan := make(chan *ProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Progress(request)
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

// ProgressWithCallback invokes the amp.Progress API asynchronously
func (client *Client) ProgressWithCallback(request *ProgressRequest, callback func(response *ProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProgressResponse
		var err error
		defer close(result)
		response, err = client.Progress(request)
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

// ProgressRequest is the request struct for api Progress
type ProgressRequest struct {
	*requests.RoaRequest
	Total     requests.Integer `position:"Body" name:"Total"`
	Test      string           `position:"Body" name:"Test"`
	ProcessId string           `position:"Body" name:"ProcessId"`
	TaskCode  string           `position:"Body" name:"TaskCode"`
	Addend    requests.Integer `position:"Body" name:"Addend"`
}

// ProgressResponse is the response struct for api Progress
type ProgressResponse struct {
	*responses.BaseResponse
}

// CreateProgressRequest creates a request to invoke Progress API
func CreateProgressRequest() (request *ProgressRequest) {
	request = &ProgressRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "Progress", "/Progress", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateProgressResponse creates a response to parse from Progress response
func CreateProgressResponse() (response *ProgressResponse) {
	response = &ProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
