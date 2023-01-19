package rest

import "github.com/fakriardian/Go-kelas.work/src/use-case/resto"

type handler struct {
	restoUseCase resto.Usecase
}

func NewHandler(restoUseCase resto.Usecase) *handler {
	return &handler{
		restoUseCase,
	}
}
