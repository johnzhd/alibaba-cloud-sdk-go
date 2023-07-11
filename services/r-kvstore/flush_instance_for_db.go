package r_kvstore

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

// FlushInstanceForDB invokes the r_kvstore.FlushInstanceForDB API synchronously
func (client *Client) FlushInstanceForDB(request *FlushInstanceForDBRequest) (response *FlushInstanceForDBResponse, err error) {
	response = CreateFlushInstanceForDBResponse()
	err = client.DoAction(request, response)
	return
}

// FlushInstanceForDBWithChan invokes the r_kvstore.FlushInstanceForDB API asynchronously
func (client *Client) FlushInstanceForDBWithChan(request *FlushInstanceForDBRequest) (<-chan *FlushInstanceForDBResponse, <-chan error) {
	responseChan := make(chan *FlushInstanceForDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FlushInstanceForDB(request)
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

// FlushInstanceForDBWithCallback invokes the r_kvstore.FlushInstanceForDB API asynchronously
func (client *Client) FlushInstanceForDBWithCallback(request *FlushInstanceForDBRequest, callback func(response *FlushInstanceForDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FlushInstanceForDBResponse
		var err error
		defer close(result)
		response, err = client.FlushInstanceForDB(request)
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

// FlushInstanceForDBRequest is the request struct for api FlushInstanceForDB
type FlushInstanceForDBRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DbIndex              requests.Integer `position:"Query" name:"DbIndex"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// FlushInstanceForDBResponse is the response struct for api FlushInstanceForDB
type FlushInstanceForDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateFlushInstanceForDBRequest creates a request to invoke FlushInstanceForDB API
func CreateFlushInstanceForDBRequest() (request *FlushInstanceForDBRequest) {
	request = &FlushInstanceForDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "FlushInstanceForDB", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFlushInstanceForDBResponse creates a response to parse from FlushInstanceForDB response
func CreateFlushInstanceForDBResponse() (response *FlushInstanceForDBResponse) {
	response = &FlushInstanceForDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
