package model

import "github.com/shopspring/decimal"

type Good struct {
	ID     uint64
	Name   string
	Price  decimal.Decimal
	Method Method
}

type Method struct{}

func (m *Method) Call() string {
	return "object method"
}

type Goods []Good
