package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"github.com/labstack/echo/v4"
)

func (h *handler) Order(c echo.Context) error {
	var request constant.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	orderData, err := h.restoUseCase.Order(request)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":   orderData,
		"status": http.StatusOK,
	})

}

func (h *handler) GetOrderInfo(c echo.Context) error {
	orderId := c.Param("orderID")

	orderData, err := h.restoUseCase.GetOrderInfo(constant.GetOrderInfoRequest{OrderID: orderId})

	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":   orderData,
		"status": http.StatusOK,
	})
}
