package resto

import (
	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
)

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	Order(request constant.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(request constant.GetOrderInfoRequest) (model.Order, error)
	RegisterUser(request constant.ResigesterUserRequest) (model.User, error)
	Login(request constant.LoginRequest) (model.UserSession, error)
}
