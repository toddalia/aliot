package iot

// 阿里云响应格式
type Response struct {
	Code string `json:"Code"`
	Success bool `json:"Success"`
	ErrorMessage string `json:"ErrorMessage"`
	RequestID string `json:"RequestId"`
	Data map[string]interface{} `json:"Data"`
}
