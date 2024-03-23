package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/veron-baranige/echo-keycloak-starter/config"
)

func GetUserGroups(params gocloak.GetGroupsParams) ([]*gocloak.Group, error) {
	return client.GetGroups(
		context.Background(),
		getClientToken().AccessToken,
		config.Env.KeycloakRealm,
		params,
	)
}