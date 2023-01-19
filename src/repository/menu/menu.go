package menu

import (
	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem

	if err := m.db.Where(model.MenuItem{Type: constant.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}
