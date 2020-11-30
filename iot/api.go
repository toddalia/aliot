package iot

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// Pub invokes aliyun `Pub` api
func Pub(client *sdk.Client, device *Device, msg *Message) (response *responses.CommonResponse, err error)  {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = fmt.Sprintf("iot.%s.aliyuncs.com", device.Region)
	request.Version = "2018-01-20"
	request.ApiName = "Pub"
	request.QueryParams["RegionId"] = device.Region
	request.QueryParams["TopicFullName"] = fmt.Sprintf("/%s/%s/user/request", device.ProductKey, device.Name)
	request.QueryParams["ProductKey"] = device.ProductKey
	content, err := msg.EncodedContent()
	if err != nil {
		return nil, err
	}
	request.QueryParams["MessageContent"] = content

	return client.ProcessCommonRequest(request)
}
