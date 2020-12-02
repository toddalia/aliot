package iot

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func buildCommonRequest(client *sdk.Client, region string) *requests.CommonRequest {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = fmt.Sprintf("iot.%s.aliyuncs.com", region)
	request.Version = "2018-01-20"
	request.QueryParams["RegionId"] = region

	return request
}

// Pub invokes aliyun `Pub` api
func Pub(client *sdk.Client, device *Device, msg *Message) (response *responses.CommonResponse, err error)  {
	request := buildCommonRequest(client, device.Region)
	request.ApiName = "Pub"
	request.QueryParams["TopicFullName"] = fmt.Sprintf("/%s/%s/user/request", device.ProductKey, device.Name)
	request.QueryParams["ProductKey"] = device.ProductKey
	content, err := msg.EncodedContent()
	if err != nil {
		return nil, err
	}
	request.QueryParams["MessageContent"] = content

	return client.ProcessCommonRequest(request)
}

// GetDeviceStatus 返回设备运行状态
func GetDeviceStatus(client *sdk.Client, device *Device) (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, device.Region)
	request.ApiName = "GetDeviceStatus"
	request.QueryParams["ProductKey"] = device.ProductKey
	request.QueryParams["DeviceName"] = device.Name

	return client.ProcessCommonRequest(request)
}

// QueryDeviceDetail 返回设备详情
func QueryDeviceDetail(client *sdk.Client, device *Device) (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, device.Region)
	request.ApiName = "QueryDeviceDetail"
	request.QueryParams["ProductKey"] = device.ProductKey
	request.QueryParams["DeviceName"] = device.Name

	return client.ProcessCommonRequest(request)
}

// QueryDeviceStatistics 返回在线设备数，设备总数
func QueryDeviceStatistics(client *sdk.Client, product *Product) (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, product.Region)
	request.ApiName = "QueryDeviceStatistics"
	request.QueryParams["ProductKey"] = product.ProductKey

	return client.ProcessCommonRequest(request)
}
