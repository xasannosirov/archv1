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
	"log"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisDB     *redis.Redis
	Enforcer    *casbin.Enforcer
	UserUseCase user.UserUseCaseI
}

func NewUserController(option *UserController) UserController {
	return UserController{
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
func (u *UserController) List(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: errStr[0],
		})
		log.Println("failed to parse query params", errStr)
		return
	}

	users, err := u.UserUseCase.List(context.Background(), entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to get list", err)
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
func (u *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind get user request", err)
		return
	}

	userResponse, err := u.UserUseCase.GetByID(context.Background(), userIntID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to get user", err)
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// Create
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
func (u *UserController) Create(c *gin.Context) {
	var request entity.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind create user request", err)
		return
	}

	userResponse, err := u.UserUseCase.Create(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to create user", err)
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}

// Update
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
func (u *UserController) Update(c *gin.Context) {
	var request entity.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind update user request", err)
		return
	}

	userResponse, err := u.UserUseCase.Update(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to update user", err)
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// UpdateColumns
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
func (u *UserController) UpdateColumns(c *gin.Context) {
	var request entity.UpdateUserColumnsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind update user request", err)
		return
	}

	userResponse, err := u.UserUseCase.UpdateColumns(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to update user", err)
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// Delete
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
func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind delete user request", err)
		return
	}

	response, err := u.UserUseCase.Delete(context.Background(), userIntID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to delete user", err)
		return
	}

	c.JSON(http.StatusOK, response)
}
