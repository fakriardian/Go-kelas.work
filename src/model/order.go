package model

import (
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
)

type Order struct {
	ID            string               `gorm:"primaryKey" json:"id"`
	UserID        string               `json:"user_id"`
	Status        constant.OrderStatus `json:"status"`
	TotalAmount   int64                `json:"total_amount"`
	ProductOrders []ProductOrder       `json:"product_orders"`
	ReferenceID   string               `gorm:"unique" json:"reference_id"`
}

type ProductOrder struct {
	ID         string                      `gorm:"primaryKey" json:"id"`
	OrderID    string                      `json:"order_id"`
	OrderCode  string                      `json:"order_code"`
	Quantity   int                         `json:"quantity"`
	TotalPrice int64                       `json:"total_price"`
	Status     constant.ProductOrderStatus `json:"status"`
}
