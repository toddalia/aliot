package iot

// Product 是阿里云物联网平台定义的产品
type Product struct {
	ProductKey string
	Region string
}

// Device 是注册在阿里云物联网平台的设备
type Device struct {
	*Product
	Name string
}
