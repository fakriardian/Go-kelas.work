package resto_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/fakriardian/Go-kelas.work/src/mocks"
	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"github.com/fakriardian/Go-kelas.work/src/use-case/resto"
)

var _ = Describe("GinkgoResto", func() {
	var useCase resto.Usecase
	var menuRepoMock *mocks.MockMenuRepository
	var orderRepoMock *mocks.MockOrderRepository
	var userRepoMock *mocks.MockUserRepository
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		menuRepoMock = mocks.NewMockMenuRepository(mockCtrl)
		orderRepoMock = mocks.NewMockOrderRepository(mockCtrl)
		userRepoMock = mocks.NewMockUserRepository(mockCtrl)

		useCase = resto.GetUseCase(menuRepoMock, orderRepoMock, userRepoMock)
	})

	Describe("request order info", func() {
		Context("it gave the correct input", func() {
			inputs := constant.GetOrderInfoRequest{
				OrderID: "valid_order_id",
				UserID:  "valid_user_id",
			}

			When("the requested orderID is not the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							ID:            "valid_order_id",
							UserID:        "valid_user_id_2",
							Status:        constant.OrderStatusFinished,
							TotalAmount:   1,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "ref_id",
						}, nil)
				})
				It("return unauthorized error", func() {
					res, err := useCase.GetOrderInfo(context.Background(), inputs)
					Expect(err).Should(HaveOccurred())
					Expect(err.Error()).To(BeEquivalentTo("unauthorized"))
					Expect(res).To(BeEquivalentTo(model.Order{}))
				})
			})

			When("the requested orderID is the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							ID:            "valid_order_id",
							UserID:        "valid_user_id",
							Status:        constant.OrderStatusFinished,
							TotalAmount:   1,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "ref_id",
						}, nil)
				})
				It("return unauthorized error", func() {
					res, err := useCase.GetOrderInfo(context.Background(), inputs)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res).To(BeEquivalentTo(model.Order{
						ID:            "valid_order_id",
						UserID:        "valid_user_id",
						Status:        constant.OrderStatusFinished,
						TotalAmount:   1,
						ProductOrders: []model.ProductOrder{},
						ReferenceID:   "ref_id",
					}))
				})
			})

		})
	})
})
