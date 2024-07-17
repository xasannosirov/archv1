package post

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/post"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type ControllerPost struct {
	Conf        *config.Config
	Postgres    *postgres.DB
	Redis       *redis.Redis
	Enforcer    *casbin.Enforcer
	PostUseCase post.PostUseCaseI
}

func NewPostController(controller *ControllerPost) *ControllerPost {
	return &ControllerPost{
		Conf:        controller.Conf,
		Postgres:    controller.Postgres,
		Redis:       controller.Redis,
		Enforcer:    controller.Enforcer,
		PostUseCase: controller.PostUseCase,
	}
}

// List
// @Summary 		List Post
// @Description 	This API for getting list post
// @Tags			post
// @Accept 			json
// @Produce 		json
// @Param 			page query int false "Page"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} entity.ListPostResponse
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/post/list [GET]
func (p *ControllerPost) List(c *gin.Context) {
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

	posts, err := p.PostUseCase.List(context.Background(), filter, lang)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetByID
// @Summary 		Get Post
// @Description 	This API for getting list with id
// @Tags			post
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Post ID"
// @Success 		200 {object} entity.GetPostResponse
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/post/{id} [GET]
func (p *ControllerPost) GetByID(c *gin.Context) {
	id := c.Param("id")

	menuIntId, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	lang := utils.GetLanguageFromHeader(c.Request)

	postResponse, err := p.PostUseCase.GetByID(context.Background(), menuIntId, lang)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, postResponse)
}

// Create
// @Security 		BearerAuth
// @Summary 		Create Post
// @Description 	This API for creating a new post
// @Tags		    post
// @Accept 			json
// @Produce 		json
// @Param 			post body entity.CreatePostRequest true "Create Post Model"
// @Success 		201 {object} entity.CreatePostResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/post [POST]
func (p *ControllerPost) Create(c *gin.Context) {
	var request entity.CreatePostRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, p.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.CreatedBy = cast.ToInt(claims["sub"])

	postResponse, err := p.PostUseCase.Create(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, postResponse)
}

// Update
// @Security 		BearerAuth
// @Summary 		Update Post
// @Description 	This API for updating a post
// @Tags		    post
// @Accept 			json
// @Produce 		json
// @Param 			post body entity.UpdatePostRequest true "Update Post Model"
// @Success 		200 {object} entity.UpdatePostResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/post [PUT]
func (p *ControllerPost) Update(c *gin.Context) {
	var request entity.UpdatePostRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, p.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.UpdatedBy = cast.ToInt(claims["sub"])

	postResponse, err := p.PostUseCase.Update(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())

		return
	}

	c.JSON(http.StatusOK, postResponse)
}

// UpdateColumns
// @Security 		BearerAuth
// @Summary 		Update Post Columns
// @Description 	This API for updating a post columns
// @Tags		    post
// @Accept 			json
// @Produce 		json
// @Param 			post body entity.UpdatePostColumnsRequest true "Update Post Columns Model"
// @Success 		200 {object} entity.UpdatePostResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/post [PATCH]
func (p *ControllerPost) UpdateColumns(c *gin.Context) {
	var request entity.UpdatePostColumnsRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, p.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.Fields["updated_by"] = cast.ToString(claims["sub"])

	postResponse, err := p.PostUseCase.UpdateColumns(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, postResponse)
}

// Delete
// @Security 		BearerAuth
// @Summary 		Delete Post
// @Description 	This API for deleting post with id
// @Tags			post
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Post ID"
// @Success 		200 {object} entity.DeletePostResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/post/{id} [DELETE]
func (p *ControllerPost) Delete(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, p.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	deletedBy := cast.ToInt(claims["sub"])

	response, err := p.PostUseCase.Delete(context.Background(), userIntID, deletedBy)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response)
}
