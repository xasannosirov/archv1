package files

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/menu"
	"archv1/internal/usecase/post"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type FileController struct {
	Conf        *config.Config
	Postgres    *postgres.DB
	Redis       *redis.Redis
	Enforcer    *casbin.Enforcer
	MenuUseCase menu.MenuUseCaseI
	PostUseCase post.PostUseCaseI
}

func NewFileController(controller *FileController) *FileController {
	return &FileController{
		Conf:        controller.Conf,
		Postgres:    controller.Postgres,
		Redis:       controller.Redis,
		Enforcer:    controller.Enforcer,
		MenuUseCase: controller.MenuUseCase,
		PostUseCase: controller.PostUseCase,
	}
}

// UploadFile
// @Summary     	Upload File
// @Security 		BearerAuth
// @Description 	This API for upload a file
// @Tags  	    	file
// @Accept      	multipart/form-data
// @Produce     	json
// @Param			file formData file true "Upload file"
// @Param 			category query string true "Category"
// @Param 			id query string true "Object ID"
// @Success     	200 {object} entity.FileUploadResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure     	500 {object} errors.Error
// @Router 			/v1/files/upload [POST]
func (f *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid file request")

		return
	}

	category := c.Query("category")
	id := c.Query("id")

	objectID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid file id")

		return
	}

	if category != "post" && category != "menu" {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid category request")

		return
	}

	uploadDir := "./internal/files"

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid mkdir request")

		return
	}

	uid := uuid.NewString()
	ext := filepath.Ext(file.Filename)

	savePath := filepath.Join(uploadDir, uid+ext)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, "error happened when save file")

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	userID := cast.ToInt(claims["sub"])

	filePath := uid + ext

	if category == "menu" {
		err = f.MenuUseCase.AddFile(context.Background(), filePath, objectID, userID)
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

			return
		}
	} else {
		err = f.PostUseCase.AddFile(context.Background(), filePath, objectID, userID)
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

			return
		}
	}

	c.JSON(http.StatusOK, entity.FileUploadResponse{
		FileURL: filePath,
	})
}

// GetFile
// @Summary 		Get File
// @Security 		BearerAuth
// @Description 	This API for getting file with url
// @Tags			file
// @Accept 			json
// @Produce 		multipart/form-data
// @Param 			url query string true "File URL"
// @Success 		200 {file} form-data "File Download"
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/files/download [GET]
func (f *FileController) GetFile(c *gin.Context) {
	baseURL := "./internal/files/"

	fileURL := c.Query("url")

	file, err := os.Open(baseURL + fileURL)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	defer file.Close()

	c.File(baseURL + fileURL)
}
