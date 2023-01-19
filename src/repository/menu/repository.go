package menu

import (
	"github.com/fakriardian/Go-kelas.work/src/model"
)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
