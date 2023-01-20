package order

import "github.com/fakriardian/Go-kelas.work/src/model"

type Reposiroty interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderID string) (model.Order, error)
}
