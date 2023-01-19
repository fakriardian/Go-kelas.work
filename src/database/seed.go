package database

import (
	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Mie Ayam",
			OrderCode: "mie_ayam",
			Price:     20000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Nasi Padang",
			OrderCode: "nasi_padang",
			Price:     18000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{{
		Name:      "Air Mineral",
		OrderCode: "air_mineral",
		Price:     5000,
		Type:      constant.MenuTypeDrink,
	}, {
		Name:      "Es Teh Manis",
		OrderCode: "es_teh_manis",
		Price:     7000,
		Type:      constant.MenuTypeDrink,
	}}

	// for migration
	// db.AutoMigrate(&model.MenuItem{})

	// for seeding
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}
