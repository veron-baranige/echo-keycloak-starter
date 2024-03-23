package auth

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/veron-baranige/echo-keycloak-starter/internal/config"
)

var (
	client *gocloak.GoCloak
)

func SetupKeycloakClient() {
	client = gocloak.NewClient(config.Get(config.KeycloakBaseURL))
	getClientToken()
}
