package oceanbasepro

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

// DescribeMetricsData invokes the oceanbasepro.DescribeMetricsData API synchronously
func (client *Client) DescribeMetricsData(request *DescribeMetricsDataRequest) (response *DescribeMetricsDataResponse, err error) {
	response = CreateDescribeMetricsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetricsDataWithChan invokes the oceanbasepro.DescribeMetricsData API asynchronously
func (client *Client) DescribeMetricsDataWithChan(request *DescribeMetricsDataRequest) (<-chan *DescribeMetricsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeMetricsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetricsData(request)
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

// DescribeMetricsDataWithCallback invokes the oceanbasepro.DescribeMetricsData API asynchronously
func (client *Client) DescribeMetricsDataWithCallback(request *DescribeMetricsDataRequest, callback func(response *DescribeMetricsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetricsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetricsData(request)
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

// DescribeMetricsDataRequest is the request struct for api DescribeMetricsData
type DescribeMetricsDataRequest struct {
	*requests.RpcRequest
	GroupByLabels string `position:"Query" name:"GroupByLabels"`
	StartTime     string `position:"Query" name:"StartTime"`
	Limit         string `position:"Query" name:"Limit"`
	SortOrder     string `position:"Query" name:"SortOrder"`
	SortMetricKey string `position:"Query" name:"SortMetricKey"`
	EndTime       string `position:"Query" name:"EndTime"`
	Labels        string `position:"Query" name:"Labels"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	Metrics       string `position:"Query" name:"Metrics"`
}

// DescribeMetricsDataResponse is the response struct for api DescribeMetricsData
type DescribeMetricsDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDescribeMetricsDataRequest creates a request to invoke DescribeMetricsData API
func CreateDescribeMetricsDataRequest() (request *DescribeMetricsDataRequest) {
	request = &DescribeMetricsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeMetricsData", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMetricsDataResponse creates a response to parse from DescribeMetricsData response
func CreateDescribeMetricsDataResponse() (response *DescribeMetricsDataResponse) {
	response = &DescribeMetricsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
