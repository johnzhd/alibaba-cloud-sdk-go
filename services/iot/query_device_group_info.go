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

// QueryDeviceGroupInfo invokes the iot.QueryDeviceGroupInfo API synchronously
func (client *Client) QueryDeviceGroupInfo(request *QueryDeviceGroupInfoRequest) (response *QueryDeviceGroupInfoResponse, err error) {
	response = CreateQueryDeviceGroupInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceGroupInfoWithChan invokes the iot.QueryDeviceGroupInfo API asynchronously
func (client *Client) QueryDeviceGroupInfoWithChan(request *QueryDeviceGroupInfoRequest) (<-chan *QueryDeviceGroupInfoResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceGroupInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceGroupInfo(request)
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

// QueryDeviceGroupInfoWithCallback invokes the iot.QueryDeviceGroupInfo API asynchronously
func (client *Client) QueryDeviceGroupInfoWithCallback(request *QueryDeviceGroupInfoRequest, callback func(response *QueryDeviceGroupInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceGroupInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceGroupInfo(request)
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

// QueryDeviceGroupInfoRequest is the request struct for api QueryDeviceGroupInfo
type QueryDeviceGroupInfoRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	GroupType         string `position:"Query" name:"GroupType"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	GroupId           string `position:"Query" name:"GroupId"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// QueryDeviceGroupInfoResponse is the response struct for api QueryDeviceGroupInfo
type QueryDeviceGroupInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryDeviceGroupInfoRequest creates a request to invoke QueryDeviceGroupInfo API
func CreateQueryDeviceGroupInfoRequest() (request *QueryDeviceGroupInfoRequest) {
	request = &QueryDeviceGroupInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceGroupInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceGroupInfoResponse creates a response to parse from QueryDeviceGroupInfo response
func CreateQueryDeviceGroupInfoResponse() (response *QueryDeviceGroupInfoResponse) {
	response = &QueryDeviceGroupInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
