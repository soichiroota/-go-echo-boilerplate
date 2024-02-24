package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {
	server.POST("/signup", Signup)
	server.POST("/login", Login)
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
