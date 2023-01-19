package main

import (
	"github.com/fakriardian/Go-kelas.work/src/database"
	"github.com/fakriardian/Go-kelas.work/src/delivery/rest"
	mRepo "github.com/fakriardian/Go-kelas.work/src/repository/menu"
	rUseCase "github.com/fakriardian/Go-kelas.work/src/use-case/resto"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database := database.GetDb(database.DbAddress)

	menuRepo := mRepo.GetRepository(database)

	restoUseCase := rUseCase.GetUseCase(menuRepo)

	handler := rest.NewHandler(restoUseCase)

	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5000"))
}
