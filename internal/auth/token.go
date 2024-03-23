package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
	"github.com/veron-baranige/echo-keycloak-starter/config"
)

var (
	clientToken *gocloak.JWT
)

func getClientToken() *gocloak.JWT {
	if clientToken != nil {
		_, claims, err := client.DecodeAccessToken(
			context.Background(),
			clientToken.AccessToken,
			config.Get(config.KeycloakRealm),
		)
		if err == nil {
			expTime, err := claims.GetExpirationTime()
			if err == nil && time.Now().Before(expTime.Time) {
				return clientToken
			}
		}
	}

	token, err := client.LoginClient(
		context.Background(),
		config.Get(config.KeycloakClientId),
		config.Get(config.KeycloakClientSecret),
		config.Get(config.KeycloakRealm),
	)
	if err != nil {
		log.Panicln("Failed to get keycloak client token:", err)
	}
	clientToken = token
	return clientToken
}

func ValidateToken(accessToken string) error {
	res, err := client.RetrospectToken(
		context.Background(),
		accessToken,
		config.Get(config.KeycloakClientId),
		config.Get(config.KeycloakClientSecret),
		config.Get(config.KeycloakRealm),
	)
	if err != nil {
		return err
	}
	if !gocloak.PBool(res.Active) {
		return fmt.Errorf("token is inactive")
	}
	return nil
}

func RefreshToken(refreshToken string) (*gocloak.JWT, error) {
	return client.RefreshToken(
		context.Background(),
		refreshToken,
		config.Get(config.KeycloakClientId),
		config.Get(config.KeycloakClientSecret),
		config.Get(config.KeycloakRealm),
	)
}

func DecodeToken(ctx context.Context,accessToken string) (*jwt.MapClaims, error) {
	_, claims, err := client.DecodeAccessToken(
		ctx,
		accessToken,
		config.Get(config.KeycloakRealm),
	)
	if err != nil {
		return nil, err
	}
	return claims, nil
}