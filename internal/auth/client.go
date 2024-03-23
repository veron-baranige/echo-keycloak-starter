package auth

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/veron-baranige/echo-keycloak-starter/config"
)

var (
	client *gocloak.GoCloak
)

func SetupKeycloakClient() {
	client = gocloak.NewClient(config.Env.KeycloakBaseURL)
	getClientToken()
}
