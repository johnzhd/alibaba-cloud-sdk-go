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

// ADClock invokes the alinlp.ADClock API synchronously
func (client *Client) ADClock(request *ADClockRequest) (response *ADClockResponse, err error) {
	response = CreateADClockResponse()
	err = client.DoAction(request, response)
	return
}

// ADClockWithChan invokes the alinlp.ADClock API asynchronously
func (client *Client) ADClockWithChan(request *ADClockRequest) (<-chan *ADClockResponse, <-chan error) {
	responseChan := make(chan *ADClockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ADClock(request)
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

// ADClockWithCallback invokes the alinlp.ADClock API asynchronously
func (client *Client) ADClockWithCallback(request *ADClockRequest, callback func(response *ADClockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ADClockResponse
		var err error
		defer close(result)
		response, err = client.ADClock(request)
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

// ADClockRequest is the request struct for api ADClock
type ADClockRequest struct {
	*requests.RpcRequest
	Params      string `position:"Body" name:"Params"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
}

// ADClockResponse is the response struct for api ADClock
type ADClockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateADClockRequest creates a request to invoke ADClock API
func CreateADClockRequest() (request *ADClockRequest) {
	request = &ADClockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "ADClock", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateADClockResponse creates a response to parse from ADClock response
func CreateADClockResponse() (response *ADClockResponse) {
	response = &ADClockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
