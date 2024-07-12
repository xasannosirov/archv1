package menu

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/menu"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControllerMenu struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisDB     *redis.Redis
	Enforcer    *casbin.Enforcer
	MenuUseCase menu.MenuUseCaseI
}

func NewMenuController(option *ControllerMenu) ControllerMenu {
	return ControllerMenu{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisDB,
		Enforcer:    option.Enforcer,
		MenuUseCase: option.MenuUseCase,
	}
}

// GetSiteMenus
// @Summary 		Get Site Menu
// @Description 	This API for getting site parent menus with children
// @Tags 			menu
// @Accept 			json
// @Produce 		json
// @Param 			page query int false "Page"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} entity.SiteMenuListResponse
// @Failure 		500 {object} errors.Error
// @Router 			/v1/site/menu/list [GET]
func (m *ControllerMenu) GetSiteMenus(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, errStr[0])

		return
	}

	lang := utils.GetLanguageFromHeader(c.Request)

	filter := entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	}

	menus, err := m.MenuUseCase.GetSiteMenus(context.Background(), filter, lang)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menus)
}

// List
// @Summary 		Get List Menu
// @Description 	This API for getting menu list
// @Tags			menu
// @Accept 			json
// @Produce 		json
// @Param 			page query int false "Page"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} entity.ListMenuResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/menu/list [GET]
func (m *ControllerMenu) List(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, errStr[0])

		return
	}

	lang := utils.GetLanguageFromHeader(c.Request)

	filter := entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	}

	menus, err := m.MenuUseCase.List(context.Background(), filter, lang)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menus)
}

// GetByID
// @Summary 		Get Menu
// @Description 	This API for getting a menu
// @Tags			menu
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Menu ID"
// @Success 		200 {object} entity.GetMenuResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/menu/{id} [GET]
func (m *ControllerMenu) GetByID(c *gin.Context) {
	id := c.Param("id")

	menuIntId, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	lang := utils.GetLanguageFromHeader(c.Request)

	menuResponse, err := m.MenuUseCase.GetByID(context.Background(), menuIntId, lang)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// Create
// @Security 		BearerAuth
// @Summary 		Create Menu
// @Description 	This API for creating a new menu
// @Tags			menu
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.CreateMenuRequest true "Create User Model"
// @Success 		201 {object} entity.CreateMenuResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/menu [POST]
func (m *ControllerMenu) Create(c *gin.Context) {
	var request entity.CreateMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, m.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.CreatedBy = claims["id"].(int)

	menuResponse, err := m.MenuUseCase.Create(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, menuResponse)
}

// Update
// @Security 		BearerAuth
// @Summary 		Update Menu
// @Description 	This API for updating a menu
// @Tags			menu
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.UpdateMenuRequest true "Update User Model"
// @Success 		200 {object} entity.UpdateMenuResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/menu [PUT]
func (m *ControllerMenu) Update(c *gin.Context) {
	var request entity.UpdateMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, m.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.UpdatedBy = claims["id"].(int)

	menuResponse, err := m.MenuUseCase.Update(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// UpdateColumns
// @Security 		BearerAuth
// @Summary 		Update Menu Columns
// @Description 	This API for updating a menu columns
// @Tags			menu
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.UpdateMenuColumnsRequest true "Update User Columns Model"
// @Success 		200 {object} entity.UpdateMenuResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/menu [PATCH]
func (m *ControllerMenu) UpdateColumns(c *gin.Context) {
	var request entity.UpdateMenuColumnsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, m.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.Fields["updated_by"] = claims["id"].(string)

	menuResponse, err := m.MenuUseCase.UpdateColumns(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// Delete
// @Security 		BearerAuth
// @Summary 		Delete menu
// @Description 	This API for deleting a menu
// @Tags			menu
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Menu ID"
// @Success 		200 {object} entity.DeleteMenuResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/menu/{id} [DELETE]
func (m *ControllerMenu) Delete(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, m.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	deletedBy := claims["id"].(int)

	response, err := m.MenuUseCase.Delete(context.Background(), userIntID, deletedBy)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response)
}
