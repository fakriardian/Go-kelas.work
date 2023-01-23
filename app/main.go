package main

import (
	"crypto/rand"
	"crypto/rsa"

	config "github.com/fakriardian/Go-kelas.work/src/database"
	"github.com/fakriardian/Go-kelas.work/src/delivery/rest"
	"github.com/fakriardian/Go-kelas.work/src/logger"
	menuRepository "github.com/fakriardian/Go-kelas.work/src/repository/menu"
	orderRepository "github.com/fakriardian/Go-kelas.work/src/repository/order"
	userRepository "github.com/fakriardian/Go-kelas.work/src/repository/user"
	rUseCase "github.com/fakriardian/Go-kelas.work/src/use-case/resto"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.Init()
	e := echo.New()

	database := config.GetDb(config.DbAddress)
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := menuRepository.GetRepository(database)
	orderRepo := orderRepository.GetRepository(database)
	userRepo, err := userRepository.GetRepository(
		database, config.Secret, config.Time, config.Memory, config.Threads, config.KeyLen, signKey, config.Exp,
	)
	if err != nil {
		panic(err)
	}

	restoUseCase := rUseCase.GetUseCase(menuRepo, orderRepo, userRepo)

	handler := rest.NewHandler(restoUseCase)

	rest.LoadMiddleware(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5000"))
}
