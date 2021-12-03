package privatelink

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

// ListVpcEndpointConnections invokes the privatelink.ListVpcEndpointConnections API synchronously
func (client *Client) ListVpcEndpointConnections(request *ListVpcEndpointConnectionsRequest) (response *ListVpcEndpointConnectionsResponse, err error) {
	response = CreateListVpcEndpointConnectionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpcEndpointConnectionsWithChan invokes the privatelink.ListVpcEndpointConnections API asynchronously
func (client *Client) ListVpcEndpointConnectionsWithChan(request *ListVpcEndpointConnectionsRequest) (<-chan *ListVpcEndpointConnectionsResponse, <-chan error) {
	responseChan := make(chan *ListVpcEndpointConnectionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpcEndpointConnections(request)
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

// ListVpcEndpointConnectionsWithCallback invokes the privatelink.ListVpcEndpointConnections API asynchronously
func (client *Client) ListVpcEndpointConnectionsWithCallback(request *ListVpcEndpointConnectionsRequest, callback func(response *ListVpcEndpointConnectionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpcEndpointConnectionsResponse
		var err error
		defer close(result)
		response, err = client.ListVpcEndpointConnections(request)
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

// ListVpcEndpointConnectionsRequest is the request struct for api ListVpcEndpointConnections
type ListVpcEndpointConnectionsRequest struct {
	*requests.RpcRequest
	EndpointId         string           `position:"Query" name:"EndpointId"`
	EndpointOwnerId    requests.Integer `position:"Query" name:"EndpointOwnerId"`
	ReplacedResourceId string           `position:"Query" name:"ReplacedResourceId"`
	NextToken          string           `position:"Query" name:"NextToken"`
	ResourceId         string           `position:"Query" name:"ResourceId"`
	ConnectionStatus   string           `position:"Query" name:"ConnectionStatus"`
	MaxResults         requests.Integer `position:"Query" name:"MaxResults"`
	EniId              string           `position:"Query" name:"EniId"`
	ServiceId          string           `position:"Query" name:"ServiceId"`
}

// ListVpcEndpointConnectionsResponse is the response struct for api ListVpcEndpointConnections
type ListVpcEndpointConnectionsResponse struct {
	*responses.BaseResponse
	NextToken   string       `json:"NextToken" xml:"NextToken"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	MaxResults  string       `json:"MaxResults" xml:"MaxResults"`
	Connections []Connection `json:"Connections" xml:"Connections"`
}

// CreateListVpcEndpointConnectionsRequest creates a request to invoke ListVpcEndpointConnections API
func CreateListVpcEndpointConnectionsRequest() (request *ListVpcEndpointConnectionsRequest) {
	request = &ListVpcEndpointConnectionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "ListVpcEndpointConnections", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVpcEndpointConnectionsResponse creates a response to parse from ListVpcEndpointConnections response
func CreateListVpcEndpointConnectionsResponse() (response *ListVpcEndpointConnectionsResponse) {
	response = &ListVpcEndpointConnectionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
