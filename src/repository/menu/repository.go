package menu

import (
	"github.com/fakriardian/Go-kelas.work/src/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
