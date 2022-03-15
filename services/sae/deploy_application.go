package sae

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

// DeployApplication invokes the sae.DeployApplication API synchronously
func (client *Client) DeployApplication(request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
	response = CreateDeployApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeployApplicationWithChan invokes the sae.DeployApplication API asynchronously
func (client *Client) DeployApplicationWithChan(request *DeployApplicationRequest) (<-chan *DeployApplicationResponse, <-chan error) {
	responseChan := make(chan *DeployApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployApplication(request)
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

// DeployApplicationWithCallback invokes the sae.DeployApplication API asynchronously
func (client *Client) DeployApplicationWithCallback(request *DeployApplicationRequest, callback func(response *DeployApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeployApplication(request)
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

// DeployApplicationRequest is the request struct for api DeployApplication
type DeployApplicationRequest struct {
	*requests.RoaRequest
	NasId                            string           `position:"Query" name:"NasId"`
	JarStartArgs                     string           `position:"Query" name:"JarStartArgs"`
	OssAkSecret                      string           `position:"Body" name:"OssAkSecret"`
	MountHost                        string           `position:"Query" name:"MountHost"`
	BatchWaitTime                    requests.Integer `position:"Query" name:"BatchWaitTime"`
	Envs                             string           `position:"Query" name:"Envs"`
	KafkaInstanceId                  string           `position:"Query" name:"KafkaInstanceId"`
	PhpPECLExtensions                string           `position:"Body" name:"PhpPECLExtensions"`
	PhpArmsConfigLocation            string           `position:"Query" name:"PhpArmsConfigLocation"`
	CustomHostAlias                  string           `position:"Query" name:"CustomHostAlias"`
	JarStartOptions                  string           `position:"Query" name:"JarStartOptions"`
	ConfigMapMountDesc               string           `position:"Body" name:"ConfigMapMountDesc"`
	OssMountDescs                    string           `position:"Body" name:"OssMountDescs"`
	KafkaEndpoint                    string           `position:"Query" name:"KafkaEndpoint"`
	PreStop                          string           `position:"Query" name:"PreStop"`
	UpdateStrategy                   string           `position:"Query" name:"UpdateStrategy"`
	ChangeOrderDesc                  string           `position:"Query" name:"ChangeOrderDesc"`
	MinReadyInstanceRatio            requests.Integer `position:"Query" name:"MinReadyInstanceRatio"`
	AutoEnableApplicationScalingRule requests.Boolean `position:"Query" name:"AutoEnableApplicationScalingRule"`
	PostStart                        string           `position:"Query" name:"PostStart"`
	PhpExtensions                    string           `position:"Body" name:"PhpExtensions"`
	AssociateEip                     requests.Boolean `position:"Body" name:"AssociateEip"`
	WebContainer                     string           `position:"Query" name:"WebContainer"`
	EnableAhas                       string           `position:"Query" name:"EnableAhas"`
	SlsConfigs                       string           `position:"Query" name:"SlsConfigs"`
	OpenCollectToKafka               requests.Boolean `position:"Query" name:"OpenCollectToKafka"`
	CommandArgs                      string           `position:"Query" name:"CommandArgs"`
	AcrAssumeRoleArn                 string           `position:"Query" name:"AcrAssumeRoleArn"`
	Readiness                        string           `position:"Query" name:"Readiness"`
	Timezone                         string           `position:"Query" name:"Timezone"`
	OssAkId                          string           `position:"Body" name:"OssAkId"`
	Liveness                         string           `position:"Query" name:"Liveness"`
	PackageVersion                   string           `position:"Query" name:"PackageVersion"`
	TomcatConfig                     string           `position:"Query" name:"TomcatConfig"`
	WarStartOptions                  string           `position:"Query" name:"WarStartOptions"`
	EdasContainerVersion             string           `position:"Query" name:"EdasContainerVersion"`
	PackageUrl                       string           `position:"Query" name:"PackageUrl"`
	TerminationGracePeriodSeconds    requests.Integer `position:"Query" name:"TerminationGracePeriodSeconds"`
	PhpConfig                        string           `position:"Body" name:"PhpConfig"`
	EnableGreyTagRoute               requests.Boolean `position:"Query" name:"EnableGreyTagRoute"`
	Command                          string           `position:"Query" name:"Command"`
	MountDesc                        string           `position:"Query" name:"MountDesc"`
	Jdk                              string           `position:"Query" name:"Jdk"`
	MinReadyInstances                requests.Integer `position:"Query" name:"MinReadyInstances"`
	KafkaLogfileConfig               string           `position:"Query" name:"KafkaLogfileConfig"`
	AcrInstanceId                    string           `position:"Body" name:"AcrInstanceId"`
	AppId                            string           `position:"Query" name:"AppId"`
	ImageUrl                         string           `position:"Query" name:"ImageUrl"`
	Php                              string           `position:"Body" name:"Php"`
	PhpConfigLocation                string           `position:"Query" name:"PhpConfigLocation"`
}

// DeployApplicationResponse is the response struct for api DeployApplication
type DeployApplicationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeployApplicationRequest creates a request to invoke DeployApplication API
func CreateDeployApplicationRequest() (request *DeployApplicationRequest) {
	request = &DeployApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DeployApplication", "/pop/v1/sam/app/deployApplication", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeployApplicationResponse creates a response to parse from DeployApplication response
func CreateDeployApplicationResponse() (response *DeployApplicationResponse) {
	response = &DeployApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
