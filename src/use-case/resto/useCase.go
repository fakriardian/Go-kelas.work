package resto

import (
	"context"

	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
)

type Usecase interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	Order(ctx context.Context, request constant.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(ctx context.Context, request constant.GetOrderInfoRequest) (model.Order, error)
	RegisterUser(ctx context.Context, request constant.ResigesterUserRequest) (model.User, error)
	Login(ctx context.Context, request constant.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}
