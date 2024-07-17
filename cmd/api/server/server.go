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

	"github.com/qori-aziz-kyc/wallet-backend/internal/config"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/routes"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/jwt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {

	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	dbUser := viper.Get("database.user")
	dbPass := viper.Get("database.pass")
	dbHost := viper.Get("database.host")
	dbPort := viper.Get("database.port")
	dbName := viper.Get("database.name")

	//setup database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// jwt service
	secretKey := viper.Get("jwt.secret").(string)
	issuer := viper.Get("jwt.issuer").(string)
	expired := viper.Get("jwt.expired").(int)
	jwtService := jwt.NewJWTService(secretKey, issuer, expired)

	// // user middleware
	// // user with valid basic token can access endpoint
	// authMiddleware := middlewares.NewAuthMiddleware(jwtService, false)

	// // admin middleware
	// // only user with valid admin token can access endpoint
	// _ = middlewares.NewAuthMiddleware(jwtService, true)

	// API Routes
	// api := router.Group("api")
	// api.GET("/", routes.RootHandler)
	// // setup router
	router := routes.SetupRouter(jwtService, db)
	// routes.NewUsersRoute(api, conn, jwtService, redisCache, ristrettoCache, authMiddleware, mailerService).Routes()

	// we can add web pages if needed
	// web := router.Group("web")
	// ...

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.Get("server.port")),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// Gracefull Shutdown
	go func() {
		fmt.Printf("success to listen and serve on :%d", 8080)
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
