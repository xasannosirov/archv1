package router

import (
	authCont "archv1/internal/controller/auth"
	chatCont "archv1/internal/controller/chat"
	fileStoreCont "archv1/internal/controller/fileStore"
	fileCont "archv1/internal/controller/files"
	menuCont "archv1/internal/controller/menu"
	postCont "archv1/internal/controller/post"
	userCont "archv1/internal/controller/user"
	_ "archv1/internal/docs"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/middleware"
	"archv1/internal/pkg/repo/cache"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/tokens"
	authRepo "archv1/internal/repository/postgres/auth"
	chatRepo "archv1/internal/repository/postgres/chat"
	fileStoreRepo "archv1/internal/repository/postgres/fileStore"
	menuRepo "archv1/internal/repository/postgres/menu"
	postRepo "archv1/internal/repository/postgres/post"
	userRepo "archv1/internal/repository/postgres/user"
	authService "archv1/internal/service/auth"
	chatService "archv1/internal/service/chat"
	fileStoreService "archv1/internal/service/fileStore"
	menuService "archv1/internal/service/menu"
	postService "archv1/internal/service/post"
	userService "archv1/internal/service/user"
	authUseCase "archv1/internal/usecase/auth"
	chatUseCase "archv1/internal/usecase/chat"
	fileStoreUseCase "archv1/internal/usecase/fileStore"
	menuUseCase "archv1/internal/usecase/menu"
	postUseCase "archv1/internal/usecase/post"
	userUseCase "archv1/internal/usecase/user"
	"archv1/internal/websocket"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Conf       *config.Config
	Hub        *websocket.Hub
	PostgresDB *postgres.DB
	RedisCache *cache.Redis
	Enforcer   *casbin.Enforcer
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

	jwtHandler := tokens.JWTHandler{
		SigningKey: option.Conf.JWTSecret,
	}

	userRepository := userRepo.NewUserRepo(option.PostgresDB)
	menuRepository := menuRepo.NewMenuRepo(option.PostgresDB)
	authRepository := authRepo.NewAuthRepo(option.PostgresDB)
	postRepository := postRepo.NewPostRepo(option.PostgresDB)
	chatRepository := chatRepo.NewChatRepo(option.PostgresDB)
	fileStoreRepository := fileStoreRepo.NewFileStoreRepo(option.PostgresDB)

	userServiceI := userService.NewUserService(userRepository)
	menuServiceI := menuService.NewMenuService(menuRepository)
	authServiceI := authService.NewAuthService(authRepository)
	postServiceI := postService.NewPostService(postRepository)
	chatServiceI := chatService.NewChatService(chatRepository)
	fileStoreServiceI := fileStoreService.NewFilesStoreService(fileStoreRepository)

	userUseCaseI := userUseCase.NewUserUseCase(userServiceI)
	menuUseCaseI := menuUseCase.NewMenuUseCase(menuServiceI)
	authUseCaseI := authUseCase.NewAuthUseCase(authServiceI)
	postUseCaseI := postUseCase.NewPostUseCase(postServiceI)
	chatUseCaseI := chatUseCase.NewChatUseCase(chatServiceI)
	fileStoreUseCaseI := fileStoreUseCase.NewFilesStoreUseCase(fileStoreServiceI)

	userController := userCont.NewUserController(&userCont.ControllerUser{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisCache,
		Enforcer:    option.Enforcer,
		UserUseCase: userUseCaseI,
		AuthUseCase: authUseCaseI,
	})

	menuController := menuCont.NewMenuController(&menuCont.ControllerMenu{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisCache,
		Enforcer:    option.Enforcer,
		MenuUseCase: menuUseCaseI,
	})

	authController := authCont.NewAuthController(&authCont.ControllerAuth{
		Conf:        option.Conf,
		PostgresDB:  option.PostgresDB,
		RedisDB:     option.RedisCache,
		Enforcer:    option.Enforcer,
		AuthUseCase: authUseCaseI,
		UserUseCase: userUseCaseI,
	})

	postController := postCont.NewPostController(&postCont.ControllerPost{
		Conf:        option.Conf,
		Postgres:    option.PostgresDB,
		Redis:       option.RedisCache,
		Enforcer:    option.Enforcer,
		PostUseCase: postUseCaseI,
	})

	fileController := fileCont.NewFileController(&fileCont.FileController{
		Conf:        option.Conf,
		Postgres:    option.PostgresDB,
		Redis:       option.RedisCache,
		Enforcer:    option.Enforcer,
		PostUseCase: postUseCaseI,
		MenuUseCase: menuUseCaseI,
	})

	filesStoreController := fileStoreCont.NewFileStoreController(fileStoreCont.ControllerFileStore{
		Conf:        option.Conf,
		Postgres:    option.PostgresDB,
		Redis:       option.RedisCache,
		Enforcer:    option.Enforcer,
		FileUseCase: fileStoreUseCaseI,
	})

	chatController := chatCont.NewChatController(&chatCont.ChatController{
		RedisCache:   option.RedisCache,
		Conf:         option.Conf,
		Hub:          option.Hub,
		Postgres:     option.PostgresDB,
		Enforcer:     option.Enforcer,
		ChatUseCaseI: chatUseCaseI,
		UserUseCase:  userUseCaseI,
	})

	router.GET("/ws", func(c *gin.Context) {
		websocket.HandleConnection(option.Hub, c.Writer, c.Request)
	})

	router.POST("/v1/auth/register", authController.Register)
	router.POST("/v1/auth/login", authController.Login)
	router.GET("/v1/auth/new-access/:refresh", authController.NewAccessToken)

	router.Use(middleware.NewAuthorizer(option.Enforcer, jwtHandler, *option.Conf))

	apiV1 := router.Group("/v1")

	// Chat APIs
	apiV1.GET("/group/user-groups/:id", chatController.UserGroups)
	apiV1.GET("/group/:id", chatController.GetGroup)
	apiV1.POST("/group", chatController.CreateGroup)
	apiV1.PUT("/group", chatController.UpdateGroup)
	apiV1.PATCH("/group", chatController.UpdateGroupColumns)
	apiV1.DELETE("/group/:id", chatController.DeleteGroup)
	apiV1.POST("/group/add-user", chatController.AddUserToGroup)
	apiV1.DELETE("/group/remove-user", chatController.RemoveUserFromGroup)
	apiV1.GET("/group/user-chats", chatController.UserChats)
	apiV1.DELETE("/group/delete-chat", chatController.DeleteChat)
	apiV1.POST("/send-message", chatController.SendMessage)
	apiV1.PUT("/update-message", chatController.UpdateMessage)
	apiV1.DELETE("/delete-message/:id", chatController.DeleteMessage)
	apiV1.GET("/chat-messages", chatController.GetChatMessages)
	apiV1.GET("/get-notifications", chatController.GetAllNotifications)
	apiV1.DELETE("/delete-chat-notifications", chatController.DeleteChatNotifications)

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

	// Upload and Download API
	apiV1.POST("/upload", fileController.UploadFile)
	apiV1.GET("/download", fileController.GetFile)

	// File Store APIs
	apiV1.GET("/folder/list", filesStoreController.ListFolder)
	apiV1.GET("/folder/:id", filesStoreController.GetFolder)
	apiV1.POST("/folder", filesStoreController.CreateFolder)
	apiV1.PUT("/folder", filesStoreController.UpdateFolder)
	apiV1.PATCH("/folder", filesStoreController.UpdateFolderColumns)
	apiV1.DELETE("/folder/:id", filesStoreController.DeleteFolder)

	apiV1.GET("/file/list", filesStoreController.ListFile)
	apiV1.GET("/file/:id", filesStoreController.GetFile)
	apiV1.POST("/file", filesStoreController.CreateFile)
	apiV1.PUT("/file", filesStoreController.UpdateFile)
	apiV1.PATCH("/file", filesStoreController.UpdateFileColumns)
	apiV1.DELETE("/file/:id", filesStoreController.DeleteFile)

	url := ginSwagger.URL("swagger/doc.json")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
