package main

import (
	"github.com/fakriardian/Go-kelas.work/src/database"
	"github.com/fakriardian/Go-kelas.work/src/delivery/rest"
	menuRepository "github.com/fakriardian/Go-kelas.work/src/repository/menu"
	orderRepository "github.com/fakriardian/Go-kelas.work/src/repository/order"
	rUseCase "github.com/fakriardian/Go-kelas.work/src/use-case/resto"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database := database.GetDb(database.DbAddress)

	menuRepo := menuRepository.GetRepository(database)
	orderRepo := orderRepository.GetRepository(database)

	restoUseCase := rUseCase.GetUseCase(menuRepo, orderRepo)

	handler := rest.NewHandler(restoUseCase)

	rest.LoadMiddleware(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5000"))
}
