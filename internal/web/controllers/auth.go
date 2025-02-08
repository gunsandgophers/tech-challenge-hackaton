package controllers

import (
	"fmt"
	"net/http"
	"tech-challenge-hackaton/internal/application/services"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type AuthController struct {
	userManagerService services.UserManagerServiceInterface
}

func NewAuthControllerController(
	userManagerService services.UserManagerServiceInterface,
) *AuthController {
	return &AuthController{
		userManagerService: userManagerService,
	}
}

type LoginRequest struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

// Checkout godoc
//
//	@Summary		Login on system
//	@Description	login at system
//	@Tags				auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body		LoginRequest	true	"Login Params"
//	@Success		200			{object}	JSONResponse{data=string}  "token"
//	@Failure		400			{object}	ErrorJSONResponse	"when bad request"
//	@Failure		406			{object}	ErrorJSONResponse	"when invalid params or invalid object"
//	@Router			/auth/login [post]
func (ac *AuthController) Login(c httpserver.HTTPContext) {
	request := LoginRequest{}
	c.BindJSON(&request)
	tokenString, err := ac.userManagerService.Login(request.User, request.Password)
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	sendSuccess(c, http.StatusOK, "login", tokenString)
}

func (ac *AuthController) EndpointProtectedByAccessToken(c httpserver.HTTPContext) {
	token, err := ac.userManagerService.ValidateAccessTokenByAuthHeader(c.GetHeader("Authorization"))
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := ac.userManagerService.GetUser(token)
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	msg := fmt.Sprintf("User ID: %s | Email: %s | Name: %s", user.ID, user.Email, user.Name)
	sendSuccess(c, http.StatusOK, "endpoint-protected-by-access-token", msg)
}
