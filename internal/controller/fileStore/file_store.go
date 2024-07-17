package fileStore

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/fileStore"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type ControllerFileStore struct {
	Conf        *config.Config
	Postgres    *postgres.DB
	Redis       *redis.Redis
	Enforcer    *casbin.Enforcer
	FileUseCase fileStore.FilesStoreUseCaseI
}

func NewFileStoreController(controller ControllerFileStore) *ControllerFileStore {
	return &ControllerFileStore{
		Conf:        controller.Conf,
		Postgres:    controller.Postgres,
		Redis:       controller.Redis,
		Enforcer:    controller.Enforcer,
		FileUseCase: controller.FileUseCase,
	}
}

// ListFolder
// @Summary 			List Folder
// Description 			This API for getting list folder
// @Tags 				folder-storage
// @Accept 				json
// @Produce 			json
// @Param 				page query int false "Page"
// @Param 				limit query int false "Limit"
// @Success 			200 {object} entity.ListFolderResponse
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/folder/list [GET]
func (f *ControllerFileStore) ListFolder(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, errStr[0])

		return
	}

	filter := entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	}

	menus, err := f.FileUseCase.ListFolder(context.Background(), filter)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menus)
}

// GetFolder
// @Summary 			Get Folder
// Description 			This API for getting a folder
// @Tags 				folder-storage
// @Accept 				json
// @Produce 			json
// @Param 				id path int true "Folder ID"
// @Success 			200 {object} entity.GetFolderResponse
// @Failure 			400 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/folder/{id} [GET]
func (f *ControllerFileStore) GetFolder(c *gin.Context) {
	id := c.Param("id")

	menuIntId, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	menuResponse, err := f.FileUseCase.GetFolder(context.Background(), menuIntId)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// CreateFolder
// @Security 			BearerAuth
// @Summary 			Create Folder
// Description 			This API for creating a new folder
// @Tags 				folder-storage
// @Accept 				json
// @Produce 			json
// @Param 				folder body entity.CreateFolderRequest true "Create Folder Model"
// @Success 			201 {object} entity.CreateFolderResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/folder [POST]
func (f *ControllerFileStore) CreateFolder(c *gin.Context) {
	var request entity.CreateFolderRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.CreatedBy = cast.ToInt(claims["sub"])

	menuResponse, err := f.FileUseCase.CreateFolder(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, menuResponse)
}

// UpdateFolder
// @Security 			BearerAuth
// @Summary 			Update Folder
// Description 			This API for updating a folder
// @Tags 				folder-storage
// @Accept 				json
// @Produce 			json
// @Param 				folder body entity.UpdateFolderRequest true "Update Folder Model"
// @Success 			200 {object} entity.UpdateFolderResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/folder [PUT]
func (f *ControllerFileStore) UpdateFolder(c *gin.Context) {
	var request entity.UpdateFolderRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.UpdatedBy = cast.ToInt(claims["sub"])

	menuResponse, err := f.FileUseCase.UpdateFolder(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// UpdateFolderColumns
// @Security 			BearerAuth
// @Summary 			Update Folder Columns
// Description 			This API for updating a folder
// @Tags 				folder-storage
// @Accept 				json
// @Produce 			json
// @Param 				folder body entity.UpdateFolderColumnsRequest true "Update Folder Columns Model"
// @Success 			200 {object} entity.UpdateFolderResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/folder [PATCH]
func (f *ControllerFileStore) UpdateFolderColumns(c *gin.Context) {
	var request entity.UpdateFolderColumnsRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.Fields["updated_by"] = cast.ToString(claims["sub"])

	menuResponse, err := f.FileUseCase.UpdateFolderColumns(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// DeleteFolder
// @Security 			BearerAuth
// @Summary 			Delete Folder
// @Description 		This API for deleting folder with id
// @Tags				folder-storage
// @Accept 				json
// @Produce 			json
// @Param 				id path int true "Folder ID"
// @Success 			200 {object} entity.DeleteFolderResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/folder/{id} [DELETE]
func (f *ControllerFileStore) DeleteFolder(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	deletedBy := cast.ToInt(claims["sub"])

	response, err := f.FileUseCase.DeleteFolder(context.Background(), userIntID, deletedBy)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response)
}

// ListFile
// @Summary 			List File
// Description 			This API for getting list file
// @Tags 				file-storage
// @Accept 				json
// @Produce 			json
// @Param 				page query int false "Page"
// @Param 				limit query int false "Limit"
// @Success 			200 {object} entity.ListFileResponse
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/file/list [GET]
func (f *ControllerFileStore) ListFile(c *gin.Context) {
	params, errStr := utils.ParseQueryParams(c.Request.URL.Query())
	if errStr != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, errStr[0])

		return
	}

	filter := entity.Filter{
		Page:  params.Page,
		Limit: params.Limit,
	}

	menus, err := f.FileUseCase.ListFile(context.Background(), filter)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menus)
}

// GetFile
// @Summary 			Get File
// Description 			This API for getting a file
// @Tags 				file-storage
// @Accept 				json
// @Produce 			json
// @Param 				id path int true "File ID"
// @Success 			200 {object} entity.GetFileResponse
// @Failure 			400 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/file/{id} [GET]
func (f *ControllerFileStore) GetFile(c *gin.Context) {
	id := c.Param("id")

	fileID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	menuResponse, err := f.FileUseCase.GetFile(context.Background(), fileID)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// CreateFile
// @Security 			BearerAuth
// @Summary 			Create File
// Description 			This API for creating a new file
// @Tags 				file-storage
// @Accept 				json
// @Produce 			json
// @Param 				folder body entity.CreateFileRequest true "Create File Model"
// @Success 			201 {object} entity.CreateFolderResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/file [POST]
func (f *ControllerFileStore) CreateFile(c *gin.Context) {
	var request entity.CreateFileRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.CreatedBy = cast.ToInt(claims["sub"])

	menuResponse, err := f.FileUseCase.CreateFile(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, menuResponse)
}

// UpdateFile
// @Security 			BearerAuth
// @Summary 			Update File
// Description 			This API for updating a file
// @Tags 				file-storage
// @Accept 				json
// @Produce 			json
// @Param 				folder body entity.UpdateFileRequest true "Update File Model"
// @Success 			200 {object} entity.UpdateFileResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/file [PUT]
func (f *ControllerFileStore) UpdateFile(c *gin.Context) {
	var request entity.UpdateFileRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.UpdatedBy = cast.ToInt(claims["sub"])

	menuResponse, err := f.FileUseCase.UpdateFile(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// UpdateFileColumns
// @Security 			BearerAuth
// @Summary 			Update File Columns
// Description 			This API for updating a file
// @Tags 				file-storage
// @Accept 				json
// @Produce 			json
// @Param 				folder body entity.UpdateFileColumnsRequest true "Update File Columns Model"
// @Success 			200 {object} entity.UpdateFileResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/file [PATCH]
func (f *ControllerFileStore) UpdateFileColumns(c *gin.Context) {
	var request entity.UpdateFileColumnsRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	request.Fields["updated_by"] = cast.ToString(claims["sub"])

	menuResponse, err := f.FileUseCase.UpdateFileColumns(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, menuResponse)
}

// DeleteFile
// @Security 			BearerAuth
// @Summary 			Delete File
// @Description 		This API for deleting file with id
// @Tags				file-storage
// @Accept 				json
// @Produce 			json
// @Param 				id path int true "File ID"
// @Success 			200 {object} entity.DeleteFileResponse
// @Failure 			400 {object} errors.Error
// @Failure 			401 {object} errors.Error
// @Failure 			403 {object} errors.Error
// @Failure 			404 {object} errors.Error
// @Failure 			500 {object} errors.Error
// @Router 				/v1/file/{id} [DELETE]
func (f *ControllerFileStore) DeleteFile(c *gin.Context) {
	id := c.Param("id")

	userIntID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, f.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	deletedBy := cast.ToInt(claims["sub"])

	response, err := f.FileUseCase.DeleteFile(context.Background(), userIntID, deletedBy)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response)
}
