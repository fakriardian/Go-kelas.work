package rest

import (
	"context"
	"net/http"

	"github.com/fakriardian/Go-kelas.work/src/tracing"
	"github.com/fakriardian/Go-kelas.work/src/use-case/resto"
	"github.com/fakriardian/Go-kelas.work/src/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func LoadMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// for multiple or com/*
		// AllowOrigins: []string{"http://restoku.com", "http://resto.com"},

		// for all
		AllowOrigins: []string{"*"},
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))
}

type authMiddleware struct {
	restoUseCase resto.Usecase
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.CreateSpan(c.Request().Context(), "CheckAuth")
		defer span.End()

		sessionData, err := utils.GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(), // for user
				Internal: err,         // for dev
			}
		}

		userID, err := am.restoUseCase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(), // for user
				Internal: err,         // for dev
			}
		}

		auhtContext := context.WithValue(c.Request().Context(), utils.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(auhtContext))

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}

func GetAuthMiddleware(restoUsecase resto.Usecase) *authMiddleware {
	return &authMiddleware{
		restoUseCase: restoUsecase,
	}
}
