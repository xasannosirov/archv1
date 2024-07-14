package auth

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/bcrypt"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/tokens"
	"archv1/internal/usecase/auth"
	"archv1/internal/usecase/user"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

type ControllerAuth struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisDB     *redis.Redis
	Enforcer    *casbin.Enforcer
	AuthUseCase auth.AuthUseCaseI
	UserUseCase user.UserUseCaseI
}

func NewAuthController(controller *ControllerAuth) ControllerAuth {
	return ControllerAuth{
		Conf:        controller.Conf,
		PostgresDB:  controller.PostgresDB,
		RedisDB:     controller.RedisDB,
		Enforcer:    controller.Enforcer,
		AuthUseCase: controller.AuthUseCase,
		UserUseCase: controller.UserUseCase,
	}
}

// Register
// @Summary 		Register
// @Description 	This API for validating email as register
// @Tags 			auth
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.RegisterRequest true "Register Model"
// @Success 		201 {object} entity.RegisterResponse
// @Failure 		400 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router			/v1/auth/register [POST]
func (a *ControllerAuth) Register(c *gin.Context) {
	var request entity.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	status, err := a.AuthUseCase.UniqueUsername(context.Background(), request.Username)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	if status {
		errors.ErrorResponse(c, http.StatusBadRequest, "Username is already taken")

		return
	}

	hashedPwd, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	userResponse, err := a.UserUseCase.Create(context.Background(), entity.CreateUserRequest{
		Username: request.Username,
		Password: hashedPwd,
		Role:     "user",
		Status:   true,
	})
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	jwtHandler := tokens.JWTHandler{
		Sub:        userResponse.Id,
		Role:       userResponse.Role,
		SigningKey: a.Conf.JWTSecret,
	}

	access, refresh, err := jwtHandler.GenerateAuthJWT()
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	err = a.AuthUseCase.UpdateToken(context.Background(), userResponse.Id, refresh)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, entity.RegisterResponse{
		ID:           userResponse.Id,
		Username:     userResponse.Username,
		Role:         userResponse.Role,
		Status:       userResponse.Status,
		AccessToken:  access,
		RefreshToken: refresh,
	})

}

// Login
// @Summary 		Login
// @Description 	This API for login
// @Tags 			auth
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.LoginRequest true "Login Model"
// @Success 		200 {object} entity.LoginResponse
// @Failure 		400 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/auth/login [POST]
func (a *ControllerAuth) Login(c *gin.Context) {
	var request entity.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	userResponse, err := a.AuthUseCase.GetUserByUsername(context.Background(), request.Username)
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())

		return
	}

	if !bcrypt.CheckPasswordHash(request.Password, userResponse.Password) {
		errors.ErrorResponse(c, http.StatusBadRequest, "Password is incorrect")

		return
	}

	jwtHandler := tokens.JWTHandler{
		Sub:        userResponse.Id,
		Role:       userResponse.Role,
		SigningKey: a.Conf.JWTSecret,
	}

	access, refresh, err := jwtHandler.GenerateAuthJWT()
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	err = a.AuthUseCase.UpdateToken(context.Background(), userResponse.Id, refresh)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, entity.LoginResponse{
		ID:           userResponse.Id,
		Username:     userResponse.Username,
		Role:         userResponse.Role,
		Status:       userResponse.Status,
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

// NewAccessToken
// @Summary 		Get New Access
// @Description 	This API for getting a new access-token
// @Tags 			auth
// @Accept 			json
// @Produce 		json
// @Param 			refresh path string true "Refresh Token"
// @Success 		200 {object} entity.NewAccessTokenResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/auth/new-access/{refresh} [GET]
func (a *ControllerAuth) NewAccessToken(c *gin.Context) {
	refreshToken := c.Param("refresh")

	userResponse, err := a.AuthUseCase.GetUserByToken(context.Background(), refreshToken)
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())

		return
	}

	claims, err := tokens.ExtractClaim(*userResponse.Refresh, []byte(a.Conf.JWTSecret))
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	userID := cast.ToInt(claims["sub"])

	jwtHandler := tokens.JWTHandler{
		Sub:        userID,
		Role:       userResponse.Role,
		SigningKey: a.Conf.JWTSecret,
	}

	access, refresh, err := jwtHandler.GenerateAuthJWT()
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	err = a.AuthUseCase.UpdateToken(context.Background(), userResponse.Id, refresh)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, entity.NewAccessTokenResponse{
		ID:           userResponse.Id,
		Username:     userResponse.Username,
		Role:         userResponse.Role,
		Status:       userResponse.Status,
		AccessToken:  access,
		RefreshToken: refresh,
	})
}
