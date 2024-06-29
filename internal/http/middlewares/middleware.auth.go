package middlewares

import (
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	jwtService jwt.JWTService
	isAdmin    bool
}

func NewAuthMiddleware(jwtService jwt.JWTService, isAdmin bool) AuthMiddleware {
	authMiddleware := AuthMiddleware{
		jwtService: jwtService,
		isAdmin:    isAdmin,
	}

	return authMiddleware
}

func (m *AuthMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "token not found")
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return handlers.NewErrorResponse(ctx, http.StatusBadRequest, "invalid header format")
		}

		if headerParts[0] != "Bearer" {
			return handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "token must content Bearer")
		}

		user, err := m.jwtService.ParseToken(headerParts[1])
		if err != nil {
			return handlers.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid token")
		}

		if user.IsAdmin != m.isAdmin && !user.IsAdmin {
			return handlers.NewErrorResponse(ctx, http.StatusForbidden, "you don't have permission to access this resource")
		}

		ctx.Set(constants.CtxAuthenticatedUserKey, user)
		return next(ctx)
	}
}
