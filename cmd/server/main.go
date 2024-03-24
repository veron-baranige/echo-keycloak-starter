package main

import (
	"log"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lmittmann/tint"
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
	setupLogger()

	if err := config.LoadEnv("."); err != nil {
		log.Fatal(err)
	}
	slog.Info("Loaded configurations")

	if err := db.SetupClient(); err != nil {
		log.Fatal(err)
	}
	slog.Info("Established database connection")
	
	auth.SetupKeycloakClient()
	slog.Info("Finished setting up keycloak client")
	
	e := echo.New()
	setupMiddleware(e)
	setupRoutes(e)
	
	if err := e.Start(":" + config.Get(config.ServerPort)); err != nil {
		slog.Error("Failed to start the server", "err", err)
	}
}

func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(m.Auth)
}

func setupRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.SetupAuthRoutes(e)
}

func setupLogger() {
	w := os.Stderr
	logger := slog.New(tint.NewHandler(w, nil))
	slog.SetDefault(logger)
}