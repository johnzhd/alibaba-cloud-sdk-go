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

// Driver is a nested struct in iot response
type Driver struct {
	DriverProtocol       string `json:"DriverProtocol" xml:"DriverProtocol"`
	IsBuiltIn            bool   `json:"IsBuiltIn" xml:"IsBuiltIn"`
	CpuArch              string `json:"CpuArch" xml:"CpuArch"`
	GmtCreateTimestamp   int64  `json:"GmtCreateTimestamp" xml:"GmtCreateTimestamp"`
	IsApply              bool   `json:"IsApply" xml:"IsApply"`
	GmtModified          string `json:"GmtModified" xml:"GmtModified"`
	Runtime              string `json:"Runtime" xml:"Runtime"`
	DriverId             string `json:"DriverId" xml:"DriverId"`
	OrderId              string `json:"OrderId" xml:"OrderId"`
	DriverVersion        string `json:"DriverVersion" xml:"DriverVersion"`
	GmtCreate            string `json:"GmtCreate" xml:"GmtCreate"`
	DriverName           string `json:"DriverName" xml:"DriverName"`
	GmtModifiedTimestamp int64  `json:"GmtModifiedTimestamp" xml:"GmtModifiedTimestamp"`
	Type                 int    `json:"Type" xml:"Type"`
}
