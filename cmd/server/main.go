package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/veron-baranige/echo-keycloak-starter/docs/swagger"
	"github.com/veron-baranige/echo-keycloak-starter/internal/auth"
	"github.com/veron-baranige/echo-keycloak-starter/internal/config"
	db "github.com/veron-baranige/echo-keycloak-starter/internal/database"
	m "github.com/veron-baranige/echo-keycloak-starter/internal/middleware"
	"github.com/veron-baranige/echo-keycloak-starter/internal/routes"
)

// @title Echo Keycloak Starter API
// @version 1.0
// @description This is a starter template for Echo with Keycloak & SQLC.
// @contact.name Veron Baranige
// @contact.email veronsajendra@gmail.com
// @host http://localhost:8080
// @BasePath /api
func main() {
	if err := config.LoadEnv("."); err != nil {
		log.Fatal(err)
	}

	if err := db.SetupClient(); err != nil {
		log.Fatal(err)
	}
	
	auth.SetupKeycloakClient()
	
	e := echo.New()
	setupMiddleware(e)
	setupRoutes(e)
	
	e.Logger.Fatal(e.Start(":" + config.Get(config.ServerPort)))
}

func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(m.Auth)
}

func setupRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.SetupAuthRoutes(e)
}