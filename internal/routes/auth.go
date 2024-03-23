package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/echo-keycloak-starter/internal/handlers"
)

func SetupAuthRoutes(e *echo.Echo) {
	auth := e.Group("/api/auth")
	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)
}