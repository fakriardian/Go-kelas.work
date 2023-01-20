package user

import "github.com/fakriardian/Go-kelas.work/src/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegister(username string) (bool, error)
	GenerateHashPassword(password string) (hash string, err error)
}
