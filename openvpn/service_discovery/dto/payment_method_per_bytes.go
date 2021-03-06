package dto

import (
	"github.com/mysterium/node/datasize"
	dto_discovery "github.com/mysterium/node/service_discovery/dto"
)

const PAYMENT_METHOD_PER_BYTES = "PER_BYTES"

type PaymentMethodPerBytes struct {
	Type string

	Price dto_discovery.Price

	// Service bytes provided for paid price
	Bytes datasize.BitSize
}

func (method PaymentMethodPerBytes) GetType() string {
	return method.Type
}

func (method PaymentMethodPerBytes) GetPrice() dto_discovery.Price {
	return method.Price
}
