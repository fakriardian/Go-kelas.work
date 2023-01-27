package order

import (
	"context"

	"github.com/fakriardian/Go-kelas.work/src/model"
)

type Reposiroty interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
}
