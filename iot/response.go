package iot

// Response 描述了阿里云响应的格式
type Response struct {
	Success bool `json:"Success"`
	Code string `json:"Code"`
	ErrorMessage string `json:"ErrorMessage"`
	RequestID string `json:"RequestId"`
	Data map[string]interface{} `json:"Data"`
}
