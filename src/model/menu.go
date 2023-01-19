package model

import (
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      constant.MenuType
}
