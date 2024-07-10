package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// setup databases
	// conn, err := utils.SetupPostgresConnection()
	// if err != nil {
	// 	return nil, err
	// }

	// setup router
	router := setupRouter()

	// // jwt service
	// jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)

	// // cache
	// redisCache := caches.NewRedisCache(config.AppConfig.REDISHost, 0, config.AppConfig.REDISPassword, time.Duration(config.AppConfig.REDISExpired))
	// ristrettoCache, err := caches.NewRistrettoCache()
	// if err != nil {
	// 	panic(err)
	// }

	// // mailer
	// mailerService := mailer.NewOTPMailer(config.AppConfig.OTPEmail, config.AppConfig.OTPPassword)

	// // user middleware
	// // user with valid basic token can access endpoint
	// authMiddleware := middlewares.NewAuthMiddleware(jwtService, false)

	// // admin middleware
	// // only user with valid admin token can access endpoint
	// _ = middlewares.NewAuthMiddleware(jwtService, true)

	// API Routes
	api := router.Group("api")
	api.GET("/", routes.RootHandler)
	// routes.NewUsersRoute(api, conn, jwtService, redisCache, ristrettoCache, authMiddleware, mailerService).Routes()

	// we can add web pages if needed
	// web := router.Group("web")
	// ...

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func setupRouter() *gin.Engine {
	// set the runtime mode
	var mode = gin.ReleaseMode
	// if config.AppConfig.Debug {
	// 	mode = gin.DebugMode
	// }
	gin.SetMode(mode)

	// create a new router instance
	router := gin.New()

	// set up middlewares
	// router.Use(middlewares.CORSMiddleware())
	// router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	return router
}

func (a *App) Run() (err error) {
	// Gracefull Shutdown
	go func() {
		fmt.Println("success to listen and serve on :%d", 8080)
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	fmt.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	fmt.Println("timeout of 5 seconds.")
	fmt.Println("server exiting")
	return
}
