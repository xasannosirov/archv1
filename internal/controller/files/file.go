package files

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/usecase/menu"
	"archv1/internal/usecase/post"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
// @Summary     Upload File
// @Description This API for upload a file
// @Tags  	    file
// @Accept      multipart/form-data
// @Produce     json
// @Param		file formData file true "Upload file"
// @Param 		request body entity.FileUploadRequest true "File Upload Model"
// @Success     200 {object} entity.FileUploadResponse
// @Failure 	400 {object} errors.Error
// @Failure 	401 {object} errors.Error
// @Failure 	403 {object} errors.Error
// @Failure     500 {object} errors.Error
// @Router 		/v1/files/upload [POST]
func (f *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid request")

		return
	}

	var request entity.FileUploadRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid request")

		return
	}

	if strings.ToLower(request.Category) != "post" || strings.ToLower(request.Category) != "menu" {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid request")

		return
	}

	uploadDir := "./files"

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid request")

		return
	}

	uid := uuid.NewString()
	ext := filepath.Ext(file.Filename)

	filePath := filepath.Join(uploadDir, uid+ext)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, "error happened when save file")

		return
	}

	if strings.ToLower(request.Category) == "menu" {
		err = f.MenuUseCase.AddFile(context.Background(), filePath, request.ObjectID)
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, "error happened when add file")

			return
		}
	} else if strings.ToLower(request.Category) == "post" {
		err = f.PostUseCase.AddFile(context.Background(), filePath, request.ObjectID)
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, "error happened when add file")

			return
		}
	}

	c.JSON(http.StatusOK, entity.FileUploadResponse{
		FileURL: filePath,
	})
}
