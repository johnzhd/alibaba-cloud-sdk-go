package config

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

// EvaluatePreConfigRules invokes the config.EvaluatePreConfigRules API synchronously
func (client *Client) EvaluatePreConfigRules(request *EvaluatePreConfigRulesRequest) (response *EvaluatePreConfigRulesResponse, err error) {
	response = CreateEvaluatePreConfigRulesResponse()
	err = client.DoAction(request, response)
	return
}

// EvaluatePreConfigRulesWithChan invokes the config.EvaluatePreConfigRules API asynchronously
func (client *Client) EvaluatePreConfigRulesWithChan(request *EvaluatePreConfigRulesRequest) (<-chan *EvaluatePreConfigRulesResponse, <-chan error) {
	responseChan := make(chan *EvaluatePreConfigRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EvaluatePreConfigRules(request)
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

// EvaluatePreConfigRulesWithCallback invokes the config.EvaluatePreConfigRules API asynchronously
func (client *Client) EvaluatePreConfigRulesWithCallback(request *EvaluatePreConfigRulesRequest, callback func(response *EvaluatePreConfigRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EvaluatePreConfigRulesResponse
		var err error
		defer close(result)
		response, err = client.EvaluatePreConfigRules(request)
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

// EvaluatePreConfigRulesRequest is the request struct for api EvaluatePreConfigRules
type EvaluatePreConfigRulesRequest struct {
	*requests.RpcRequest
	EnableManagedRules    requests.Boolean `position:"Body" name:"EnableManagedRules"`
	ResourceEvaluateItems string           `position:"Body" name:"ResourceEvaluateItems"`
	ResourceTypeFormat    string           `position:"Body" name:"ResourceTypeFormat"`
}

// EvaluatePreConfigRulesResponse is the response struct for api EvaluatePreConfigRules
type EvaluatePreConfigRulesResponse struct {
	*responses.BaseResponse
	RequestId           string               `json:"RequestId" xml:"RequestId"`
	ResourceEvaluations []ResourceEvaluation `json:"ResourceEvaluations" xml:"ResourceEvaluations"`
}

// CreateEvaluatePreConfigRulesRequest creates a request to invoke EvaluatePreConfigRules API
func CreateEvaluatePreConfigRulesRequest() (request *EvaluatePreConfigRulesRequest) {
	request = &EvaluatePreConfigRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "EvaluatePreConfigRules", "", "")
	request.Method = requests.POST
	return
}

// CreateEvaluatePreConfigRulesResponse creates a response to parse from EvaluatePreConfigRules response
func CreateEvaluatePreConfigRulesResponse() (response *EvaluatePreConfigRulesResponse) {
	response = &EvaluatePreConfigRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
