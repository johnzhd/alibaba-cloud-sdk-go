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

// QueryDeviceOriginalServiceData invokes the iot.QueryDeviceOriginalServiceData API synchronously
func (client *Client) QueryDeviceOriginalServiceData(request *QueryDeviceOriginalServiceDataRequest) (response *QueryDeviceOriginalServiceDataResponse, err error) {
	response = CreateQueryDeviceOriginalServiceDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceOriginalServiceDataWithChan invokes the iot.QueryDeviceOriginalServiceData API asynchronously
func (client *Client) QueryDeviceOriginalServiceDataWithChan(request *QueryDeviceOriginalServiceDataRequest) (<-chan *QueryDeviceOriginalServiceDataResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceOriginalServiceDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceOriginalServiceData(request)
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

// QueryDeviceOriginalServiceDataWithCallback invokes the iot.QueryDeviceOriginalServiceData API asynchronously
func (client *Client) QueryDeviceOriginalServiceDataWithCallback(request *QueryDeviceOriginalServiceDataRequest, callback func(response *QueryDeviceOriginalServiceDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceOriginalServiceDataResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceOriginalServiceData(request)
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

// QueryDeviceOriginalServiceDataRequest is the request struct for api QueryDeviceOriginalServiceData
type QueryDeviceOriginalServiceDataRequest struct {
	*requests.RpcRequest
	NextPageToken string           `position:"Query" name:"NextPageToken"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Identifier    string           `position:"Query" name:"Identifier"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	Asc           requests.Integer `position:"Query" name:"Asc"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// QueryDeviceOriginalServiceDataResponse is the response struct for api QueryDeviceOriginalServiceData
type QueryDeviceOriginalServiceDataResponse struct {
	*responses.BaseResponse
	RequestId    string                               `json:"RequestId" xml:"RequestId"`
	Success      bool                                 `json:"Success" xml:"Success"`
	Code         string                               `json:"Code" xml:"Code"`
	ErrorMessage string                               `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceOriginalServiceData `json:"Data" xml:"Data"`
}

// CreateQueryDeviceOriginalServiceDataRequest creates a request to invoke QueryDeviceOriginalServiceData API
func CreateQueryDeviceOriginalServiceDataRequest() (request *QueryDeviceOriginalServiceDataRequest) {
	request = &QueryDeviceOriginalServiceDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceOriginalServiceData", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceOriginalServiceDataResponse creates a response to parse from QueryDeviceOriginalServiceData response
func CreateQueryDeviceOriginalServiceDataResponse() (response *QueryDeviceOriginalServiceDataResponse) {
	response = &QueryDeviceOriginalServiceDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
