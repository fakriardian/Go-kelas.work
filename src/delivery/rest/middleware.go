package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// for multiple or com/*
		// AllowOrigins: []string{"http://restoku.com", "http://resto.com"},

		// for all
		AllowOrigins: []string{"*"},
	}))
}
