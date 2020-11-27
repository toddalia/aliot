package iot

type Product struct {
	ProductKey string
	Region string
}

type Device struct {
	Product
	Name string
}
