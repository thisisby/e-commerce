package main

import (
	"fmt"
	"ga_marketplace/internal/config"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/http/middlewares"
	"ga_marketplace/internal/http/routes"
	"ga_marketplace/internal/utils"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/mobizon"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func main() {
	conn, err := utils.SetupPostgreConnection()
	if err != nil {
		slog.Error("failed to create app: ", err)
		return
	}
	defer conn.Close()

	slog.Info("success to connect to database")

	// cache
	redisHost := fmt.Sprintf("%s:%d", config.AppConfig.RedisHost, config.AppConfig.RedisPort)
	redisCache := caches.NewRedisCache(redisHost, 0, config.AppConfig.RedisPassword, time.Duration(config.AppConfig.RedisExpires))

	// mobizon integration
	mobizonClient := mobizon.NewClient(config.AppConfig.MobizonBaseUrl, config.AppConfig.MobizonApiKey)

	// jwt
	jwtService := jwt.NewJWTService()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, host=${host}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Validator = utils.NewValidator()

	clientAuthMiddleware := middlewares.NewAuthMiddleware(jwtService, false)
	adminAuthMiddleware := middlewares.NewAuthMiddleware(jwtService, true)

	v1 := e.Group("/api/v1")
	routes.NewRolesRoute(conn, v1).Register()
	routes.NewUsersRoute(conn, v1, redisCache, jwtService, clientAuthMiddleware, adminAuthMiddleware, mobizonClient).Register()

	slog.Info("success to listen and serve on :8080")
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.AppConfig.Port)))

}
