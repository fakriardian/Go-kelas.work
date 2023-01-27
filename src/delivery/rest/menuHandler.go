package rest

import (
	"net/http"

	"github.com/fakriardian/Go-kelas.work/src/tracing"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) GetMenuList(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetMenuList")
	defer span.End()

	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUseCase.GetMenuList(ctx, menuType)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][handler][GetMenuList] failed get menu list")
		// fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	logrus.WithFields(logrus.Fields{
		"status": http.StatusOK,
	}).Error("[delivery][rest][handler][GetMenuList] success Get menu list")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":   menuData,
		"status": http.StatusOK,
	})
}
