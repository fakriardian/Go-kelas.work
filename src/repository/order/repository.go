package order

import (
	"context"

	"github.com/fakriardian/Go-kelas.work/src/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockOrderRepository -destination=../../mocks/orderRepositoryMock.go -source=repository.go

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
}
