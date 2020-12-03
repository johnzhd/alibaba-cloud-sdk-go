package qualitycheck

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

// UpdateSubScoreForApi invokes the qualitycheck.UpdateSubScoreForApi API synchronously
func (client *Client) UpdateSubScoreForApi(request *UpdateSubScoreForApiRequest) (response *UpdateSubScoreForApiResponse, err error) {
	response = CreateUpdateSubScoreForApiResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSubScoreForApiWithChan invokes the qualitycheck.UpdateSubScoreForApi API asynchronously
func (client *Client) UpdateSubScoreForApiWithChan(request *UpdateSubScoreForApiRequest) (<-chan *UpdateSubScoreForApiResponse, <-chan error) {
	responseChan := make(chan *UpdateSubScoreForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSubScoreForApi(request)
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

// UpdateSubScoreForApiWithCallback invokes the qualitycheck.UpdateSubScoreForApi API asynchronously
func (client *Client) UpdateSubScoreForApiWithCallback(request *UpdateSubScoreForApiRequest, callback func(response *UpdateSubScoreForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSubScoreForApiResponse
		var err error
		defer close(result)
		response, err = client.UpdateSubScoreForApi(request)
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

// UpdateSubScoreForApiRequest is the request struct for api UpdateSubScoreForApi
type UpdateSubScoreForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UpdateSubScoreForApiResponse is the response struct for api UpdateSubScoreForApi
type UpdateSubScoreForApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateSubScoreForApiRequest creates a request to invoke UpdateSubScoreForApi API
func CreateUpdateSubScoreForApiRequest() (request *UpdateSubScoreForApiRequest) {
	request = &UpdateSubScoreForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UpdateSubScoreForApi", "Qualitycheck", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSubScoreForApiResponse creates a response to parse from UpdateSubScoreForApi response
func CreateUpdateSubScoreForApiResponse() (response *UpdateSubScoreForApiResponse) {
	response = &UpdateSubScoreForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
