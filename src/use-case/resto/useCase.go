package resto

import "github.com/fakriardian/Go-kelas.work/src/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
