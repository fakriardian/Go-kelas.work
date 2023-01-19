package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
	dbAddress     = "host=localhost port=5432 user=postgres password=password dbname=go_resto_app sslmode=disable"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Mie Ayam",
			OrderCode: "mie_ayam",
			Price:     20000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Nasi Padang",
			OrderCode: "nasi_padang",
			Price:     18000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{{
		Name:      "Air Mineral",
		OrderCode: "air_mineral",
		Price:     5000,
		Type:      MenuTypeDrink,
	}, {
		Name:      "Es Teh Manis",
		OrderCode: "es_teh_manis",
		Price:     7000,
		Type:      MenuTypeDrink,
	}}

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	// for migration
	// db.AutoMigrate(&MenuItem{})

	// for seeding
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}

// func getFoodMenu(c echo.Context) error {

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		// "data": foodMenu,
// 	})
// }

// func getDrinkMenu(c echo.Context) error {

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": http.StatusOK,
// 		// "data":   drinkMenu,
// 	})
// }

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   menuData,
	})
}

func main() {
	seedDB()
	e := echo.New()
	// route
	// e.GET("/menu/food", getFoodMenu)
	// e.GET("/menu/drink", getDrinkMenu)
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":5000"))
}
