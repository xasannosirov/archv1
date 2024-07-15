package fileStore

import (
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/usecase/fileStore"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
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

func (f *ControllerFileStore) ListFolder(c *gin.Context) {}

func (f *ControllerFileStore) GetFolder(c *gin.Context) {}

func (f *ControllerFileStore) CreateFolder(c *gin.Context) {}

func (f *ControllerFileStore) UpdateFolder(c *gin.Context) {}

func (f *ControllerFileStore) UpdateFolderColumns(c *gin.Context) {}

func (f *ControllerFileStore) DeleteFolder(c *gin.Context) {}

func (f *ControllerFileStore) ListFile(c *gin.Context) {}

func (f *ControllerFileStore) GetFile(c *gin.Context) {}

func (f *ControllerFileStore) CreateFile(c *gin.Context) {}

func (f *ControllerFileStore) UpdateFile(c *gin.Context) {}

func (f *ControllerFileStore) UpdateFileColumns(c *gin.Context) {}

func (f *ControllerFileStore) DeleteFile(c *gin.Context) {}
