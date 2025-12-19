package model

type Stock struct {
	UUID     string `json:"uuid"`
	StockId  string `json:"stock_id" validate:"required"`
	Name     string `json:"string" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
	Company  string `json:"company" validate:"required"`
}
