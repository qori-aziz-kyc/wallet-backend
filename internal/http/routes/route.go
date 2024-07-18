package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/handlers/roothandler"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/middlewares"
	"github.com/qori-aziz-kyc/wallet-backend/internal/injection"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/jwt"
	"gorm.io/gorm"
)

func SetupRouter(jwt jwt.JWTService, db *gorm.DB) *gin.Engine {
	// set the runtime mode
	var mode = gin.ReleaseMode
	gin.SetMode(mode)

	handler := injection.NewInitialInjection(jwt, db)

	// create a new router instance
	router := gin.New()

	// set up middlewares
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// API Routes
	api := router.Group("api")
	api.GET("/", roothandler.RootHandler)

	adminMiddlerware := middlewares.NewAuthMiddleware(jwt, true)
	userMiddlerware := middlewares.NewAuthMiddleware(jwt, false)

	{
		api.POST("register", handler.User.RegisterHandler)
		api.POST("login", handler.User.LoginHandler)
	}

	{
		categoryAdminAPI := api.Group("categories")
		categoryAdminAPI.Use(adminMiddlerware)
		categoryAdminAPI.POST("/", handler.Category.CreateHandler)
		categoryAdminAPI.PUT("/:id", handler.Category.UpdateHandler)

		categoryUserAPI := api.Group("categories")
		categoryUserAPI.Use(userMiddlerware)
		categoryUserAPI.GET("/", handler.Category.FindHandler)
	}

	{
		recordAPI := api.Group("records")
		recordAPI.Use(userMiddlerware)
		recordAPI.POST("/", handler.Record.CreateHandler)
		recordAPI.PUT("/:id", handler.Record.UpdateHandler)
		recordAPI.GET("/", handler.Record.FindHandler)
	}

	for _, item := range router.Routes() {
		fmt.Println(item.Method, item.Path)
	}

	return router
}
