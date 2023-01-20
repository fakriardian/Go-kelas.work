package resto

import (
	"errors"

	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"github.com/fakriardian/Go-kelas.work/src/repository/menu"
	"github.com/fakriardian/Go-kelas.work/src/repository/order"
	"github.com/fakriardian/Go-kelas.work/src/repository/user"
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

func (r *restoUseCase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUseCase) Order(request constant.OrderMenuRequest) (model.Order, error) {
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
		Status:        constant.OrderStatusProcessed,
		TotalAmount:   amounTotal,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createOrderData, nil

}

func (r *restoUseCase) GetOrderInfo(request constant.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderID)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

func (r *restoUseCase) RegisterUser(request constant.ResigesterUserRequest) (model.User, error) {
	userRegistered, err := r.userRepo.CheckRegister(request.UserName)
	if err != nil {
		return model.User{}, err
	}

	if userRegistered {
		return model.User{}, errors.New("already registered!")
	}

	passwordHash, err := r.userRepo.GenerateHashPassword(request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(model.User{
		ID:       uuid.NewString(),
		UserName: request.UserName,
		Password: passwordHash,
	})

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}
