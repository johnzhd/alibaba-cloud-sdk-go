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

// SyncSpeechByCombination invokes the iot.SyncSpeechByCombination API synchronously
func (client *Client) SyncSpeechByCombination(request *SyncSpeechByCombinationRequest) (response *SyncSpeechByCombinationResponse, err error) {
	response = CreateSyncSpeechByCombinationResponse()
	err = client.DoAction(request, response)
	return
}

// SyncSpeechByCombinationWithChan invokes the iot.SyncSpeechByCombination API asynchronously
func (client *Client) SyncSpeechByCombinationWithChan(request *SyncSpeechByCombinationRequest) (<-chan *SyncSpeechByCombinationResponse, <-chan error) {
	responseChan := make(chan *SyncSpeechByCombinationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncSpeechByCombination(request)
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

// SyncSpeechByCombinationWithCallback invokes the iot.SyncSpeechByCombination API asynchronously
func (client *Client) SyncSpeechByCombinationWithCallback(request *SyncSpeechByCombinationRequest, callback func(response *SyncSpeechByCombinationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncSpeechByCombinationResponse
		var err error
		defer close(result)
		response, err = client.SyncSpeechByCombination(request)
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

// SyncSpeechByCombinationRequest is the request struct for api SyncSpeechByCombination
type SyncSpeechByCombinationRequest struct {
	*requests.RpcRequest
	SpeechId        string    `position:"Body" name:"SpeechId"`
	AudioFormat     string    `position:"Body" name:"AudioFormat"`
	IotId           string    `position:"Body" name:"IotId"`
	CombinationList *[]string `position:"Body" name:"CombinationList"  type:"Repeated"`
	IotInstanceId   string    `position:"Body" name:"IotInstanceId"`
	ProductKey      string    `position:"Body" name:"ProductKey"`
	ApiProduct      string    `position:"Body" name:"ApiProduct"`
	ApiRevision     string    `position:"Body" name:"ApiRevision"`
	DeviceName      string    `position:"Body" name:"DeviceName"`
}

// SyncSpeechByCombinationResponse is the response struct for api SyncSpeechByCombination
type SyncSpeechByCombinationResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInSyncSpeechByCombination `json:"Data" xml:"Data"`
}

// CreateSyncSpeechByCombinationRequest creates a request to invoke SyncSpeechByCombination API
func CreateSyncSpeechByCombinationRequest() (request *SyncSpeechByCombinationRequest) {
	request = &SyncSpeechByCombinationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SyncSpeechByCombination", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSyncSpeechByCombinationResponse creates a response to parse from SyncSpeechByCombination response
func CreateSyncSpeechByCombinationResponse() (response *SyncSpeechByCombinationResponse) {
	response = &SyncSpeechByCombinationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
