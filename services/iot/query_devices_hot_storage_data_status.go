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

// QueryDevicesHotStorageDataStatus invokes the iot.QueryDevicesHotStorageDataStatus API synchronously
func (client *Client) QueryDevicesHotStorageDataStatus(request *QueryDevicesHotStorageDataStatusRequest) (response *QueryDevicesHotStorageDataStatusResponse, err error) {
	response = CreateQueryDevicesHotStorageDataStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDevicesHotStorageDataStatusWithChan invokes the iot.QueryDevicesHotStorageDataStatus API asynchronously
func (client *Client) QueryDevicesHotStorageDataStatusWithChan(request *QueryDevicesHotStorageDataStatusRequest) (<-chan *QueryDevicesHotStorageDataStatusResponse, <-chan error) {
	responseChan := make(chan *QueryDevicesHotStorageDataStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDevicesHotStorageDataStatus(request)
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

// QueryDevicesHotStorageDataStatusWithCallback invokes the iot.QueryDevicesHotStorageDataStatus API asynchronously
func (client *Client) QueryDevicesHotStorageDataStatusWithCallback(request *QueryDevicesHotStorageDataStatusRequest, callback func(response *QueryDevicesHotStorageDataStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDevicesHotStorageDataStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryDevicesHotStorageDataStatus(request)
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

// QueryDevicesHotStorageDataStatusRequest is the request struct for api QueryDevicesHotStorageDataStatus
type QueryDevicesHotStorageDataStatusRequest struct {
	*requests.RpcRequest
	NextPageToken     string           `position:"Query" name:"NextPageToken"`
	RealTenantId      string           `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string           `position:"Query" name:"RealTripartiteKey"`
	IotId             string           `position:"Query" name:"IotId"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	UserTopic         string           `position:"Query" name:"UserTopic"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	Asc               requests.Integer `position:"Query" name:"Asc"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	DeviceName        string           `position:"Query" name:"DeviceName"`
}

// QueryDevicesHotStorageDataStatusResponse is the response struct for api QueryDevicesHotStorageDataStatus
type QueryDevicesHotStorageDataStatusResponse struct {
	*responses.BaseResponse
	RequestId    string                                 `json:"RequestId" xml:"RequestId"`
	Success      bool                                   `json:"Success" xml:"Success"`
	Code         string                                 `json:"Code" xml:"Code"`
	ErrorMessage string                                 `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDevicesHotStorageDataStatus `json:"Data" xml:"Data"`
}

// CreateQueryDevicesHotStorageDataStatusRequest creates a request to invoke QueryDevicesHotStorageDataStatus API
func CreateQueryDevicesHotStorageDataStatusRequest() (request *QueryDevicesHotStorageDataStatusRequest) {
	request = &QueryDevicesHotStorageDataStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDevicesHotStorageDataStatus", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDevicesHotStorageDataStatusResponse creates a response to parse from QueryDevicesHotStorageDataStatus response
func CreateQueryDevicesHotStorageDataStatusResponse() (response *QueryDevicesHotStorageDataStatusResponse) {
	response = &QueryDevicesHotStorageDataStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
