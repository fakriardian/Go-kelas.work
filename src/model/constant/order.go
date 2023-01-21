package constant

type OrderStatus string

type ProductOrderStatus string

const (
	OrderStatusProcessed OrderStatus = "processed"
	OrderStatusFinished  OrderStatus = "finished"
	OrderStatusFailed    OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing ProductOrderStatus = "preparing"
	ProductOrderStatusFinished  ProductOrderStatus = "finished"
)

// like schema
type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	UserID        string                    `json:"-"`
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	UserID  string `json:"-"`
	OrderID string `json:"order_id"`
}
