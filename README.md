# Echo Starter with Keycloak and SQLC

## Dependencies
- Make: `sudo apt install make`
- SQLC: `sudo snap install sqlc`
- Air: `go install github.com/cosmtrek/air@latest`
- Swagger: `go get -u github.com/swaggo/swag`
- Golang Migrate:  
  ```
  curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
  apt-get update
  apt-get install -y migrate
  ```

## Used Packages
- Viper: `go get github.com/spf13/viper`
- GoCloak: `go get github.com/Nerzal/gocloak/v13`
- Echo: `go get github.com/labstack/echo/v4`
- UUID: `go get github.com/google/uuid`
- Testify: `go get github.com/stretchr/testify`
- EchoSwagger: `go get -u github.com/swaggo/echo-swagger`

## Usage
### Server
- Setup dev environment with databases and keycloak
  ```docker compose -f deployments/docker-compose.yml up -d```
- Start dev server with live reloading: `air`

### Configurations
- Change the environment variable in the __.env__ file according to your dev environment

### Database
- Create migrate file: `make migrate-init name={NAME}`
- Migrate up: `make migrate-up`
- Migrate down: `make migrate-down`
- Migrate force fix: `make migrate-fix v={VERSION}`

## Swagger
- Annotate handlers with valid swagger comments
- Docs will be auto generate on each modification
- [Visit swagger](http://localhost:8000/swagger/index.html)