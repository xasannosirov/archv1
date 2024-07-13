package router

import (
	authControllerPackage "archv1/internal/controller/auth"
	fileControllerPackage "archv1/internal/controller/files"
	menuControllerPackage "archv1/internal/controller/menu"
	postControllerPackage "archv1/internal/controller/post"
	userControllerPackage "archv1/internal/controller/user"
	_ "archv1/internal/docs"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	authRepoPackage "archv1/internal/repository/postgres/auth"
	menuRepoPackage "archv1/internal/repository/postgres/menu"
	postRepoPackage "archv1/internal/repository/postgres/post"
	userRepoPackage "archv1/internal/repository/postgres/user"
	authServicePackage "archv1/internal/service/auth"
	menuServicePackage "archv1/internal/service/menu"
	postServicePackage "archv1/internal/service/post"
	userServicePackage "archv1/internal/service/user"
	authUseCasePackage "archv1/internal/usecase/auth"
	menuUseCasePackage "archv1/internal/usecase/menu"
	postUseCasePackage "archv1/internal/usecase/post"
	userUseCasePackage "archv1/internal/usecase/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Conf        *config.Config
	PostgresDB  *postgres.DB
	RedisCache  *redis.Redis
	Enforcer    *casbin.Enforcer
	UserUseCase userUseCasePackage.UserUseCaseI
}

// New
// @securityDefinitions.apikey BearerAuth
// @In          	header
// @Name        	Authorization
func New(option *Router) *gin.Engine {
	gin.SetMode(option.Conf.GinMode)
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//jwtHandler := tokens.JWTHandler{
	//	SigningKey: option.Conf.JWTSecret,
	//}

	//middleware.NewAuthorizer(option.Enforcer, jwtHandler, *option.Conf)

	apiV1 := router.Group("/v1")

	userRepository := userRepoPackage.NewUserRepo(option.PostgresDB)
	menuRepository := menuRepoPackage.NewMenuRepo(option.PostgresDB)
	authRepository := authRepoPackage.NewAuthRepo(option.PostgresDB)
	postRepository := postRepoPackage.NewPostRepo(option.PostgresDB)

	userService := userServicePackage.NewUserService(userRepository)
	menuService := menuServicePackage.NewMenuService(menuRepository)
	authService := authServicePackage.NewAuthService(authRepository)
	postService := postServicePackage.NewPostService(postRepository)

	userUseCase := userUseCasePackage.NewUserUseCase(userService)
	menuUseCase := menuUseCasePackage.NewMenuUseCase(menuService)
	authUseCase := authUseCasePackage.NewAuthUseCase(authService)
	postUseCase := postUseCasePackage.NewPostUseCase(postService)

	userController := userControllerPackage.NewUserController(&userControllerPackage.ControllerUser{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisCache,
		Enforcer:    option.Enforcer,
		UserUseCase: userUseCase,
	})

	menuController := menuControllerPackage.NewMenuController(&menuControllerPackage.ControllerMenu{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisCache,
		Enforcer:    option.Enforcer,
		MenuUseCase: menuUseCase,
	})

	authController := authControllerPackage.NewAuthController(&authControllerPackage.ControllerAuth{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisCache,
		Enforcer:    option.Enforcer,
		AuthUseCase: authUseCase,
	})

	postController := postControllerPackage.NewPostController(&postControllerPackage.ControllerPost{
		Conf:        option.Conf,
		Postgres:    option.PostgresDB,
		Redis:       option.RedisCache,
		Enforcer:    option.Enforcer,
		PostUseCase: postUseCase,
	})

	fileController := fileControllerPackage.NewFileController(&fileControllerPackage.FileController{
		Conf:        option.Conf,
		Postgres:    option.PostgresDB,
		Redis:       option.RedisCache,
		Enforcer:    option.Enforcer,
		PostUseCase: postUseCase,
	})

	// Auth APIs
	apiV1.POST("/auth/register", authController.Register)
	apiV1.POST("/auth/login", authController.Login)
	apiV1.GET("/auth/new-access/:refresh", authController.NewAccessToken)

	// User APIs
	apiV1.GET("/user/list", userController.List)
	apiV1.GET("/user/:id", userController.GetByID)
	apiV1.POST("/user", userController.Create)
	apiV1.PUT("/user", userController.Update)
	apiV1.PATCH("/user", userController.UpdateColumns)
	apiV1.DELETE("/user/:id", userController.Delete)

	// Menu APIs
	apiV1.GET("/site/menu/list", menuController.GetSiteMenus)
	apiV1.GET("/menu/list", menuController.List)
	apiV1.GET("/menu/:id", menuController.GetByID)
	apiV1.POST("/menu", menuController.Create)
	apiV1.PUT("/menu", menuController.Update)
	apiV1.PATCH("/menu", menuController.UpdateColumns)
	apiV1.DELETE("/menu/:id", menuController.Delete)

	// Post APIs
	apiV1.GET("/post/list", postController.List)
	apiV1.GET("/post/:id", postController.GetByID)
	apiV1.POST("/post", postController.Create)
	apiV1.PUT("/post", postController.Update)
	apiV1.PATCH("/post", postController.UpdateColumns)
	apiV1.DELETE("/post/:id", postController.Delete)

	// File APIs
	apiV1.POST("/upload-file", fileController.UploadFile)

	url := ginSwagger.URL("swagger/doc.json")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
