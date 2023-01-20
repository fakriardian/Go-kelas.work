package resto

import (
	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
)

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	Order(request constant.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(request constant.GetOrderInfoRequest) (model.Order, error)
}
