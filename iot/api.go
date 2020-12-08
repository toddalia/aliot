package iot

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func buildCommonRequest(client *sdk.Client, product *Product) *requests.CommonRequest {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Version = "2018-01-20"
	request.Domain = fmt.Sprintf("iot.%s.aliyuncs.com", product.Region)
	request.QueryParams["RegionId"] = product.Region
	request.QueryParams["ProductKey"] = product.ProductKey

	return request
}

// Pub 调用阿里云 `Pub` 接口，发送文本消息
func Pub(client *sdk.Client, device *Device, msg string)  (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, device.Product)
	request.ApiName = "Pub"
	request.QueryParams["TopicFullName"] = fmt.Sprintf("/%s/%s/user/request", device.ProductKey, device.Name)
	request.QueryParams["MessageContent"] = msg

	return client.ProcessCommonRequest(request)
}

// PubMessage 调用阿里云 `Pub` 接口，向设备发送自定义格式消息
func PubMessage(client *sdk.Client, device *Device, msg *Message) (response *responses.CommonResponse, err error)  {
	content, err := msg.EncodedContent()
	if err != nil {
		return nil, err
	}
	return Pub(client, device, content)
}

// GetDeviceStatus 返回设备运行状态
func GetDeviceStatus(client *sdk.Client, device *Device) (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, device.Product)
	request.ApiName = "GetDeviceStatus"
	request.QueryParams["DeviceName"] = device.Name

	return client.ProcessCommonRequest(request)
}

// QueryDeviceDetail 返回设备详情
func QueryDeviceDetail(client *sdk.Client, device *Device) (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, device.Product)
	request.ApiName = "QueryDeviceDetail"
	request.QueryParams["DeviceName"] = device.Name

	return client.ProcessCommonRequest(request)
}

// QueryDeviceStatistics 返回在线设备数，设备总数
func QueryDeviceStatistics(client *sdk.Client, product *Product) (response *responses.CommonResponse, err error) {
	request := buildCommonRequest(client, product)
	request.ApiName = "QueryDeviceStatistics"

	return client.ProcessCommonRequest(request)
}
