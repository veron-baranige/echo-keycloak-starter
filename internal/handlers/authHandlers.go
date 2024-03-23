package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/echo-keycloak-starter/internal/auth"
	db "github.com/veron-baranige/echo-keycloak-starter/internal/database"
)

type (
	UserRegistrationRequest struct {
		FirstName    string `json:"firstName"`
		LastName     string `json:"lasttName"`
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}

	UserLoginRequest struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}

	UserLoginResponse struct {
		AccessToken  string       `json:"accessToken"`
		RefreshToken string       `json:"refreshToken"`
		User         LoginUserDto `json:"user"`
	}
	LoginUserDto struct {
		FirstName    string       `json:"firstName"`
		LastName     string       `json:"lastName"`
		EmailAddress string       `json:"emailAddress"`
		Role         db.UsersRole `json:"role"`
	}
)

// @Summary Create new user account
// @Accept  json
// @Param object body UserRegistrationRequest true "User Registration Request"
// @Success 201 
// @Failure 400
// @Failure 500
// @Router /auth/register [post]
// @Tags Auth
func RegisterUser(c echo.Context) error {
	var userReq UserRegistrationRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&userReq); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	_, err := db.Client.GetUserByEmail(context.Background(), userReq.EmailAddress)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return echo.ErrBadRequest
	}

	keycloakUid, err := auth.RegisterUser(userReq.EmailAddress, userReq.Password, string(db.UsersRoleUSER))
	if err != nil {
		log.Println("Failed to create keycloak user: ", err)
		return echo.ErrServiceUnavailable
	}

	_, err = db.Client.CreateUser(context.Background(), db.CreateUserParams{
		ID:           uuid.NewString(),
		FirstName:    userReq.FirstName,
		LastName:     userReq.LastName,
		EmailAddress: userReq.EmailAddress,
		KeycloakUid:  keycloakUid,
		Role:         db.UsersRoleUSER,
	})
	if err != nil {
		log.Println(err)
		if err := auth.RemoveKeycloakUser(keycloakUid); err != nil {
			log.Println("Failed to remove keycloak user: ", err)
		}
		return echo.ErrInternalServerError
	}

	return c.String(http.StatusCreated, "Created user successfully")
}

// @Summary User login
// @Accept json
// @Produce json 
// @Param object body UserLoginRequest true "User Login Request"
// @Success 200 {object} UserLoginResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /auth/login [post]
// @Tags Auth
func LoginUser(c echo.Context) error {
	var loginReq UserLoginRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&loginReq); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	jwt, err := auth.LoginUser(loginReq.EmailAddress, loginReq.Password)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	user, err := db.Client.GetUserByEmail(context.Background(), loginReq.EmailAddress)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, UserLoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		User: LoginUserDto{
			FirstName: user.FirstName,
			LastName: user.LastName,
			EmailAddress: user.EmailAddress,
			Role: user.Role,
		},
	})
}
