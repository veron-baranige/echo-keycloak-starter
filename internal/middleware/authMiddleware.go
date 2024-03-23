package middleware

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/echo-keycloak-starter/internal/auth"
	db "github.com/veron-baranige/echo-keycloak-starter/internal/database"
)

var (
	publicRotes = map[string]bool{
		"/api/auth/register":      true,
		"/api/auth/login":         true,
	}
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().URL.String(), "/swagger") {
			return next(c)
		}

		_, ok := publicRotes[c.Request().URL.String()]
		if ok {
			return next(c)
		}

		authHeaders, ok := c.Request().Header["Authorization"]
		if !ok || len(authHeaders) == 0 || !strings.Contains(authHeaders[0], "Bearer") {
			return echo.ErrUnauthorized
		}

		token := strings.ReplaceAll(strings.ReplaceAll(authHeaders[0], "Bearer", ""), " ", "")
		if token == "" {
			return echo.ErrUnauthorized
		}

		if err := auth.ValidateToken(token); err != nil {
			return echo.ErrUnauthorized
		}

		claims, err := auth.DecodeToken(context.Background(), token)
		if err != nil || (*claims)["sub"] == "" {
			return echo.ErrUnauthorized
		}

		keycloakUid, ok := (*claims)["sub"].(string)
		if !ok {
			return echo.ErrUnauthorized
		}

		user, err := db.Client.GetActiveKeycloakUser(context.Background(), keycloakUid)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Set("user", user)

		perms, ok := (*claims)["realm_access"].(map[string]interface{})["roles"].([]interface{})
		if ok {
			c.Set("permissions", perms)
		}

		return next(c)
	}
}
