package auth

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/usecase/auth"
	"archv1/internal/usecase/user"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthController struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisDB     *redis.Redis
	Enforcer    *casbin.Enforcer
	AuthUseCase auth.AuthUseCaseI
	UserUseCase user.UserUseCaseI
}

func NewAuthController(controller *AuthController) AuthController {
	return AuthController{
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
func (a *AuthController) Register(c *gin.Context) {
	var request entity.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind request", err)
		return
	}

	status, err := a.AuthUseCase.UniqueUsername(context.Background(), request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to check unique username", err)
		return
	}
	if status {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: "Username already taken",
		})
		log.Println("username already used", err)
		return
	}

	// hashed password

	userResponse, err := a.UserUseCase.Create(context.Background(), entity.CreateUserRequest{
		Username: request.Username,
		Password: request.Password,
		Role:     "user",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to create user", err)
		return
	}

	// generate token

	c.JSON(http.StatusCreated, entity.RegisterResponse{
		ID:           userResponse.Id,
		Username:     userResponse.Username,
		Role:         userResponse.Role,
		Status:       true,
		AccessToken:  "",
		RefreshToken: "",
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
func (a *AuthController) Login(c *gin.Context) {}

// NewAccessToken
// @Summary 		Get New Access
// @Description 	This API for getting a new access-token
// @Tags 			auth
// @Accept 			json
// @Produce 		json
// @Param 			refresh path string true "Refresh Token"
// @Success 		200 {object} entity.NewAccessTokenResponse
// @Failure 		400 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/auth/new-access/:refresh [GET]
func (a *AuthController) NewAccessToken(c *gin.Context) {}
