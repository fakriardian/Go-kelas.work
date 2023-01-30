package menu

import (
	"context"

	"github.com/fakriardian/Go-kelas.work/src/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockMenuRepository -destination=../../mocks/menuRepositoryMock.go -source=repository.go

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
