package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/veron-baranige/echo-keycloak-starter/config"
	db "github.com/veron-baranige/echo-keycloak-starter/database"
	"github.com/veron-baranige/echo-keycloak-starter/internal/auth"
	m "github.com/veron-baranige/echo-keycloak-starter/internal/middleware"
	"github.com/veron-baranige/echo-keycloak-starter/internal/routes"
)

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
	
	e.Logger.Fatal(e.Start(":" + config.Env.ServerPort))
}

func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(m.Auth)
}

func setupRoutes(e *echo.Echo) {
	routes.SetupAuthRoutes(e)
}