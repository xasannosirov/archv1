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
	"log"
	"net/http"
	"strconv"
)

type MenuController struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisDB     *redis.Redis
	Enforcer    *casbin.Enforcer
	MenuUseCase menu.MenuUseCaseI
}

func NewMenuController(option *MenuController) MenuController {
	return MenuController{
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
func (m *MenuController) GetSiteMenus(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: errStr[0],
		})
		log.Println("failed to parse query params", errStr)
		return
	}

	lang := c.Request.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}

	filter := entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	}

	menus, err := m.MenuUseCase.GetSiteMenus(context.Background(), filter, lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to get parent menus with child", err)
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
func (m *MenuController) List(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: errStr[0],
		})
		log.Println("failed to parse query params", errStr)
		return
	}

	lang := c.Request.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}

	filter := entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	}

	menus, err := m.MenuUseCase.List(context.Background(), filter, lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to get list menu", err)
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
func (m *MenuController) GetByID(c *gin.Context) {
	id := c.Param("id")

	menuIntId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind get menu request", err)
		return
	}

	lang := c.Request.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}

	menuResponse, err := m.MenuUseCase.GetByID(context.Background(), menuIntId, lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to get menu", err)
		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// Create
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
func (m *MenuController) Create(c *gin.Context) {
	var request entity.CreateMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind create menu request", err)
		return
	}

	menuResponse, err := m.MenuUseCase.Create(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to create menu", err)
		return
	}

	c.JSON(http.StatusCreated, menuResponse)
}

// Update
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
func (m *MenuController) Update(c *gin.Context) {
	var request entity.UpdateMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind update menu request", err)
		return
	}

	menuResponse, err := m.MenuUseCase.Update(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to update menu", err)
		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// UpdateColumns
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
func (m *MenuController) UpdateColumns(c *gin.Context) {
	var request entity.UpdateMenuColumnsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind update menu request", err)
		return
	}

	menuResponse, err := m.MenuUseCase.UpdateColumns(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to update menu", err)
		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// Delete
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
func (m *MenuController) Delete(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to bind delete menu request", err)
		return
	}

	response, err := m.MenuUseCase.Delete(context.Background(), userIntID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error{
			Message: err.Error(),
		})
		log.Println("failed to delete menu", err)
		return
	}

	c.JSON(http.StatusOK, response)
}
