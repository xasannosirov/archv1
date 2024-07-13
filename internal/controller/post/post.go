package post

import (
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/usecase/post"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
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

func (p *ControllerPost) List(c *gin.Context) {}

func (p *ControllerPost) GetByID(c *gin.Context) {}

func (p *ControllerPost) Create(c *gin.Context) {}

func (p *ControllerPost) Update(c *gin.Context) {}

func (p *ControllerPost) UpdateColumns(c *gin.Context) {}

func (p *ControllerPost) Delete(c *gin.Context) {}
