package resto

import (
	"context"
	"errors"

	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"github.com/fakriardian/Go-kelas.work/src/repository/menu"
	"github.com/fakriardian/Go-kelas.work/src/repository/order"
	"github.com/fakriardian/Go-kelas.work/src/repository/user"
	"github.com/fakriardian/Go-kelas.work/src/tracing"
	"github.com/google/uuid"
)

type restoUseCase struct {
	menuRepo  menu.Repository
	orderRepo order.Reposiroty
	userRepo  user.Repository
}

func GetUseCase(menuRepo menu.Repository, orderRepo order.Reposiroty, userRepo user.Repository) Usecase {
	return &restoUseCase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restoUseCase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	return r.menuRepo.GetMenuList(ctx, menuType)
}

func (r *restoUseCase) Order(ctx context.Context, request constant.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()

	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))
	var amounTotal int64 = 0

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		amounTotal += (menuData.Price * int64(orderProduct.Quantity))

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menuData.Price * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.NewString(),
		UserID:        request.UserID,
		Status:        constant.OrderStatusProcessed,
		TotalAmount:   amounTotal,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createOrderData, err := r.orderRepo.CreateOrder(ctx, orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createOrderData, nil

}

func (r *restoUseCase) GetOrderInfo(ctx context.Context, request constant.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()

	orderData, err := r.orderRepo.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		return orderData, err
	}

	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("unauthorized")
	}

	return orderData, nil
}

func (r *restoUseCase) RegisterUser(ctx context.Context, request constant.ResigesterUserRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	userRegistered, err := r.userRepo.CheckRegister(ctx, request.UserName)
	if err != nil {
		return model.User{}, err
	}

	if userRegistered {
		return model.User{}, errors.New("already registered!")
	}

	passwordHash, err := r.userRepo.GenerateHashPassword(ctx, request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(ctx, model.User{
		ID:       uuid.NewString(),
		UserName: request.UserName,
		Password: passwordHash,
	})

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (r *restoUseCase) Login(ctx context.Context, request constant.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()

	userData, err := r.userRepo.GetUserData(ctx, request.UserName)
	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := r.userRepo.VerifyLogin(ctx, request.UserName, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}

	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := r.userRepo.CreateUserSession(ctx, userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}

func (r *restoUseCase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	userID, err = r.userRepo.CheckSession(ctx, data)
	if err != nil {
		return "", err
	}

	return userID, nil

}
