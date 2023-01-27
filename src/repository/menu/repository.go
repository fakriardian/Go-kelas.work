package menu

import (
	"context"

	"github.com/fakriardian/Go-kelas.work/src/model"
)

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
