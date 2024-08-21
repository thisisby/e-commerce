package jwt

import (
	"errors"
	"ga_marketplace/internal/config"
	"ga_marketplace/pkg/helpers"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTService interface {
	GenerateToken(userId int, isAdmin bool, duration time.Duration) (string, error)
	ParseToken(tokenString string) (JWTCustomClaims, error)
}

type JWTCustomClaims struct {
	UserId  int
	IsAdmin bool
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	secretKey, issuer := getConfigClaims()
	return &jwtService{
		secretKey,
		issuer,
	}
}

func (j *jwtService) GenerateToken(userId int, isAdmin bool, duration time.Duration) (string, error) {
	claims := JWTCustomClaims{
		userId,
		isAdmin,
		jwt.StandardClaims{
			Issuer:   j.issuer,
			IssuedAt: helpers.GetCurrentTime().Unix(),
		},
	}

	if !isAdmin {
		claims.ExpiresAt = helpers.GetCurrentTime().Add(time.Hour * duration).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtService) ParseToken(tokenString string) (JWTCustomClaims, error) {
	var claims JWTCustomClaims

	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	}); err != nil || !token.Valid {
		return JWTCustomClaims{}, errors.New("invalid token")
	}

	return claims, nil
}

func getConfigClaims() (secretKey, issuer string) {
	issuer = config.AppConfig.JWTIssuer
	secretKey = config.AppConfig.JWTSecret

	return
}
