package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/veron-baranige/echo-keycloak-starter/config"
)

func RegisterUser(emailAddress string, password string, role string) (uid string, err error) {
	return client.CreateUser(
		context.Background(),
		getClientToken().AccessToken,
		config.Get(config.KeycloakRealm),
		gocloak.User{
			Username:      gocloak.StringP(emailAddress),
			Email:         gocloak.StringP(emailAddress),
			EmailVerified: gocloak.BoolP(false),
			Enabled:       gocloak.BoolP(true),
			Credentials: &[]gocloak.CredentialRepresentation{
				{
					Temporary: gocloak.BoolP(false),
					Type:      gocloak.StringP(password),
					Value:     gocloak.StringP(password),
				},
			},
			Groups: &[]string{role},
		},
	)
}

func LoginUser(emailAddress string, password string) (*gocloak.JWT, error) {
	return client.Login(
		context.Background(),
		config.Get(config.KeycloakClientId),
		config.Get(config.KeycloakClientSecret),
		config.Get(config.KeycloakRealm),
		emailAddress,
		password,
	)
}

func LogoutUser(uid string) error {
	return client.LogoutAllSessions(
		context.Background(),
		getClientToken().AccessToken,
		config.Get(config.KeycloakRealm),
		uid,
	)
}

func RemoveKeycloakUser(uid string) error {
	return client.DeleteUser(
		context.Background(),
		getClientToken().AccessToken,
        config.Get(config.KeycloakRealm),
        uid,
	)
}