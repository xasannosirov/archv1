package user_controller

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/user"
	"context"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisDB     *redis.Redis
	Enforcer    *casbin.Enforcer
	UserUseCase user.UserUseCaseI
}

func NewUserController(option *ControllerUser) ControllerUser {
	return ControllerUser{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisDB,
		Enforcer:    option.Enforcer,
		UserUseCase: option.UserUseCase,
	}
}

// List
// @Summary 		Get List User
// @Description 	This API for getting user list
// @Tags			user
// @Accept 			json
// @Produce 		json
// @Param 			page query int false "Page"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} entity.ListUserResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/user/list [GET]
func (u *ControllerUser) List(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, errStr[0])

		return
	}

	users, err := u.UserUseCase.List(context.Background(), entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	})
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, users)
}

// GetByID
// @Summary 		Get User
// @Description 	This API for getting a user
// @Tags			user
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} entity.GetUserResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/user/{id} [GET]
func (u *ControllerUser) GetByID(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	userResponse, err := u.UserUseCase.GetByID(context.Background(), userIntID)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// Create
// @Security 		BearerAuth
// @Summary 		Create User
// @Description 	This API for creating a new user
// @Tags			user
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.CreateUserRequest true "Create User Model"
// @Success 		201 {object} entity.CreateUserResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/user [POST]
func (u *ControllerUser) Create(c *gin.Context) {
	var request entity.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, u.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.CreatedBy = cast.ToInt(claims["sub"])

	userRole := strings.ToLower(request.Role)
	if userRole != "user" && userRole != "admin" {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid role")

		return
	}

	userResponse, err := u.UserUseCase.Create(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, userResponse)
}

// Update
// @Security 		BearerAuth
// @Summary 		Update User
// @Description 	This API for updating a user
// @Tags			user
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.UpdateUserRequest true "Update User Model"
// @Success 		200 {object} entity.UpdateUserResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/user [PUT]
func (u *ControllerUser) Update(c *gin.Context) {
	var request entity.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	userRole := strings.ToLower(request.Role)
	if userRole != "user" && userRole != "admin" {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid role")

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, u.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.UpdatedBy = cast.ToInt(claims["sub"])

	userResponse, err := u.UserUseCase.Update(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// UpdateColumns
// @Security 		BearerAuth
// @Summary 		Update User Columns
// @Description 	This API for updating a user columns
// @Tags			user
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.UpdateUserColumnsRequest true "Update User Columns Model"
// @Success 		200 {object} entity.UpdateUserResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/user [PATCH]
func (u *ControllerUser) UpdateColumns(c *gin.Context) {
	var request entity.UpdateUserColumnsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if request.Fields["role"] != "" {
		userRole := strings.ToLower(request.Fields["role"])
		if userRole != "user" && userRole != "admin" {
			errors.ErrorResponse(c, http.StatusBadRequest, "invalid role")

			return
		}
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, u.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.Fields["updated_by"] = cast.ToString(claims["sub"])

	userResponse, err := u.UserUseCase.UpdateColumns(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// Delete
// @Security 		BearerAuth
// @Summary 		Delete User
// @Description 	This API for deleting a user
// @Tags			user
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} entity.DeleteUserResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/user/{id} [DELETE]
func (u *ControllerUser) Delete(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, u.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	deletedBy := cast.ToInt(claims["sub"])

	response, err := u.UserUseCase.Delete(context.Background(), userIntID, deletedBy)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response)
}
