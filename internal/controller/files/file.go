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
// @Summary     Upload File
// @Description This API for upload a file
// @Tags  	    file
// @Accept      multipart/form-data
// @Produce     json
// @Param		file formData file true "Upload file"
// @Param 		category query string true "Category"
// @Param 		id query string true "Object ID"
// @Success     200 {object} entity.FileUploadResponse
// @Failure 	400 {object} errors.Error
// @Failure 	401 {object} errors.Error
// @Failure 	403 {object} errors.Error
// @Failure     500 {object} errors.Error
// @Router 		/v1/files/upload [POST]
func (f *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, "invalid file request")

		return
	}

	category := c.Query("category")
	id := c.Query("id")

	intId, err := strconv.Atoi(id)
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
	addr := "localhost:8080/"

	savePath := filepath.Join(uploadDir, uid+ext)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, "error happened when save file")

		return
	}

	filePath := addr + savePath

	if category == "menu" {
		err = f.MenuUseCase.AddFile(context.Background(), filePath, intId)
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

			return
		}
	} else {
		err = f.PostUseCase.AddFile(context.Background(), filePath, intId)
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

			return
		}
	}

	c.JSON(http.StatusOK, entity.FileUploadResponse{
		FileURL: filePath,
	})
}
