package ecs

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

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
	response = CreateDescribeSnapshotsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSnapshotsWithChan(request *DescribeSnapshotsRequest) (<-chan *DescribeSnapshotsResponse, <-chan error) {
	responseChan := make(chan *DescribeSnapshotsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnapshots(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeSnapshotsWithCallback(request *DescribeSnapshotsRequest, callback func(response *DescribeSnapshotsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnapshotsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnapshots(request)
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

type DescribeSnapshotsRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	SourceDiskType       string           `position:"Query" name:"SourceDiskType"`
	Filter2Key           string           `position:"Query" name:"Filter.2.Key"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
	Filter2Value         string           `position:"Query" name:"Filter.2.Value"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Filter1Value         string           `position:"Query" name:"Filter.1.Value"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	SnapshotType         string           `position:"Query" name:"SnapshotType"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SnapshotIds          string           `position:"Query" name:"SnapshotIds"`
	Status               string           `position:"Query" name:"Status"`
	SnapshotName         string           `position:"Query" name:"SnapshotName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	SnapshotLinkId       string           `position:"Query" name:"SnapshotLinkId"`
	Encrypted            requests.Boolean `position:"Query" name:"Encrypted"`
	Filter1Key           string           `position:"Query" name:"Filter.1.Key"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	DiskId               string           `position:"Query" name:"DiskId"`
	Usage                string           `position:"Query" name:"Usage"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
}

type DescribeSnapshotsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Snapshots  struct {
		Snapshot []struct {
			SnapshotId        string `json:"SnapshotId" xml:"SnapshotId"`
			SnapshotName      string `json:"SnapshotName" xml:"SnapshotName"`
			Progress          string `json:"Progress" xml:"Progress"`
			ProductCode       string `json:"ProductCode" xml:"ProductCode"`
			SourceDiskId      string `json:"SourceDiskId" xml:"SourceDiskId"`
			SourceDiskType    string `json:"SourceDiskType" xml:"SourceDiskType"`
			RetentionDays     int    `json:"RetentionDays" xml:"RetentionDays"`
			Encrypted         bool   `json:"Encrypted" xml:"Encrypted"`
			SourceDiskSize    int    `json:"SourceDiskSize" xml:"SourceDiskSize"`
			Description       string `json:"Description" xml:"Description"`
			CreationTime      string `json:"CreationTime" xml:"CreationTime"`
			Status            string `json:"Status" xml:"Status"`
			Usage             string `json:"Usage" xml:"Usage"`
			SourceStorageType string `json:"SourceStorageType" xml:"SourceStorageType"`
			Tags              struct {
				Tag []struct {
					TagKey   string `json:"TagKey" xml:"TagKey"`
					TagValue string `json:"TagValue" xml:"TagValue"`
				} `json:"Tag" xml:"Tag"`
			} `json:"Tags" xml:"Tags"`
		} `json:"Snapshot" xml:"Snapshot"`
	} `json:"Snapshots" xml:"Snapshots"`
}

func CreateDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshots", "", "")
	return
}

func CreateDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
