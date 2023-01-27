package rest

import (
	"encoding/json"
	"net/http"

	"github.com/fakriardian/Go-kelas.work/src/model/constant"
	"github.com/fakriardian/Go-kelas.work/src/tracing"
	"github.com/fakriardian/Go-kelas.work/src/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()

	var request constant.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][handler][Order] request failed")
		// fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	userID := c.Request().Context().Value(utils.AuthContextKey).(string)
	request.UserID = userID

	orderData, err := h.restoUseCase.Order(ctx, request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][handler][Order] failed order data")
		// fmt.Printf("got error %s\n", err.Error())

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
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetOrderInfo")
	defer span.End()

	orderId := c.Param("orderID")
	userID := c.Request().Context().Value(utils.AuthContextKey).(string)

	orderData, err := h.restoUseCase.GetOrderInfo(ctx, constant.GetOrderInfoRequest{
		UserID:  userID,
		OrderID: orderId,
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][handler][GetOrderInfo] unable to get order data")

		// fmt.Printf("got error %s\n", err.Error())

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
