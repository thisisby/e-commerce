package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/mobizon"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type UsersRoute struct {
	usersHandler   handlers.UsersHandler
	router         *echo.Group
	db             *sqlx.DB
	authMiddleware middlewares.AuthMiddleware
	adminMid       middlewares.AuthMiddleware
	mobizonClient  mobizon.Client
}

func NewUsersRoute(
	db *sqlx.DB,
	router *echo.Group,
	redisCache caches.RedisCache,
	jwtService jwt.JWTService,
	clientAuthMiddleware middlewares.AuthMiddleware,
	authMiddleware middlewares.AuthMiddleware,
	mobizonClient mobizon.Client,
) *UsersRoute {
	userRepo := postgre.NewPostgreUsersRepository(db)
	userUsecase := usecases.NewUsersUsecase(userRepo, jwtService, mobizonClient)
	usersHandler := handlers.NewUsersHandler(userUsecase, redisCache)

	return &UsersRoute{
		usersHandler:   usersHandler,
		router:         router,
		db:             db,
		authMiddleware: clientAuthMiddleware,
		adminMid:       authMiddleware,
		mobizonClient:  mobizonClient,
	}
}

func (r *UsersRoute) Register() {
	//auth
	auth := r.router.Group("/auth")
	auth.POST("/send-otp", r.usersHandler.SendOTP)
	auth.POST("/reset-attempts", r.usersHandler.ResetAttempts)
	auth.POST("/verify-otp", r.usersHandler.VerifyOTP)
	auth.POST("/register", r.usersHandler.Register)
	auth.GET("/refresh-token", r.usersHandler.RefreshToken)
	auth.POST("/resend-otp", r.usersHandler.ResendOTP)

	auth.Use(r.authMiddleware.Handle)
	auth.GET("/me", r.usersHandler.GetMe)

	//users
	users := r.router.Group("/users")
	admin := r.router.Group("/admin/users")

	users.Use(r.authMiddleware.Handle)
	users.PATCH("", r.usersHandler.UpdateMe)
	users.DELETE("", r.usersHandler.DeleteMe)

	admin.Use(r.adminMid.Handle)
	admin.GET("", r.usersHandler.GetAllUsers)
	admin.DELETE("/:id", r.usersHandler.DeleteUser)

}
