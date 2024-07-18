package main

import (
	"fmt"
	_ "ga_marketplace/docs"
	"ga_marketplace/internal/config"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/http/middlewares"
	"ga_marketplace/internal/http/routes"
	"ga_marketplace/internal/utils"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/aws"
	"ga_marketplace/third_party/mobizon"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Logger initialized")

	if err := config.InitializeAppConfig(); err != nil {
		slog.Error("failed to initialize app config: ", err)
		return
	}
}

//	@title			GA Marketplace API
//	@version		1.0
//	@description	This is a sample server with null types overridden with primitive types.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		product_info.swagger.io
//	@BasePath	/v2

func main() {
	conn, err := utils.SetupPostgreConnection()
	if err != nil {
		slog.Error("failed to create app: ", err)
		return
	}

	slog.Info("success to connect to database")

	// cache
	redisHost := fmt.Sprintf("%s:%d", config.AppConfig.RedisHost, config.AppConfig.RedisPort)
	redisCache := caches.NewRedisCache(redisHost, 0, config.AppConfig.RedisPassword, time.Duration(config.AppConfig.RedisExpires))

	// mobizon integration
	mobizonClient := mobizon.NewClient(config.AppConfig.MobizonBaseUrl, config.AppConfig.MobizonApiKey)

	// s3 integration
	s3Client := aws.NewS3Client()
	buckets := s3Client.ListBuckets()
	fmt.Printf("buckets: %v\n", buckets)

	// jwt
	jwtService := jwt.NewJWTService()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, host=${host}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Validator = utils.NewValidator()

	clientAuthMiddleware := middlewares.NewAuthMiddleware(jwtService, false)
	adminAuthMiddleware := middlewares.NewAuthMiddleware(jwtService, true)

	v1 := e.Group("/api/v1")
	routes.NewRolesRoute(conn, v1).Register()
	routes.NewUsersRoute(conn, v1, redisCache, jwtService, clientAuthMiddleware, adminAuthMiddleware, mobizonClient).Register()
	routes.NewCartsRoute(conn, v1, redisCache, clientAuthMiddleware, adminAuthMiddleware).Register()
	routes.NewWishRoute(conn, v1, clientAuthMiddleware).Register()
	routes.NewDiscountRoute(conn, v1, adminAuthMiddleware).Register()
	routes.NewProfileSectionsRoute(conn, v1, clientAuthMiddleware, adminAuthMiddleware).Register()
	routes.NewProductRoute(conn, v1, s3Client, clientAuthMiddleware, adminAuthMiddleware).Register()
	routes.NewHealthCheckRoute(v1).Register()
	routes.NewCountriesRoute(conn, v1, clientAuthMiddleware, adminAuthMiddleware).Register()
	routes.NewOrdersRoute(conn, v1, clientAuthMiddleware).Register()

	slog.Info("success to listen and serve on :8080")

	slog.Info("buckets: ", buckets)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.AppConfig.Port)))

}
