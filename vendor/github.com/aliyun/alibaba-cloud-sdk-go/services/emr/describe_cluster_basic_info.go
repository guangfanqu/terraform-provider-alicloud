package emr

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

// DescribeClusterBasicInfo invokes the emr.DescribeClusterBasicInfo API synchronously
// api document: https://help.aliyun.com/api/emr/describeclusterbasicinfo.html
func (client *Client) DescribeClusterBasicInfo(request *DescribeClusterBasicInfoRequest) (response *DescribeClusterBasicInfoResponse, err error) {
	response = CreateDescribeClusterBasicInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterBasicInfoWithChan invokes the emr.DescribeClusterBasicInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/describeclusterbasicinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterBasicInfoWithChan(request *DescribeClusterBasicInfoRequest) (<-chan *DescribeClusterBasicInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterBasicInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterBasicInfo(request)
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

// DescribeClusterBasicInfoWithCallback invokes the emr.DescribeClusterBasicInfo API asynchronously
// api document: https://help.aliyun.com/api/emr/describeclusterbasicinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterBasicInfoWithCallback(request *DescribeClusterBasicInfoRequest, callback func(response *DescribeClusterBasicInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterBasicInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterBasicInfo(request)
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

// DescribeClusterBasicInfoRequest is the request struct for api DescribeClusterBasicInfo
type DescribeClusterBasicInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// DescribeClusterBasicInfoResponse is the response struct for api DescribeClusterBasicInfo
type DescribeClusterBasicInfoResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ClusterInfo ClusterInfo `json:"ClusterInfo" xml:"ClusterInfo"`
}

// CreateDescribeClusterBasicInfoRequest creates a request to invoke DescribeClusterBasicInfo API
func CreateDescribeClusterBasicInfoRequest() (request *DescribeClusterBasicInfoRequest) {
	request = &DescribeClusterBasicInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterBasicInfo", "emr", "openAPI")
	return
}

// CreateDescribeClusterBasicInfoResponse creates a response to parse from DescribeClusterBasicInfo response
func CreateDescribeClusterBasicInfoResponse() (response *DescribeClusterBasicInfoResponse) {
	response = &DescribeClusterBasicInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}