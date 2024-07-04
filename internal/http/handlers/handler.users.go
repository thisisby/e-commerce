package handlers

import (
	"errors"
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
)

type UsersHandler struct {
	userUsecase domains.UserUsecase
	redisCache  caches.RedisCache
}

func NewUsersHandler(userUsecase domains.UserUsecase, redisCache caches.RedisCache) UsersHandler {
	return UsersHandler{
		userUsecase: userUsecase,
		redisCache:  redisCache,
	}
}

func (u *UsersHandler) SendOTP(ctx echo.Context) error {
	var userSendOTP requests.UserSendOTPRequest

	err := helpers.BindAndValidate(ctx, &userSendOTP)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	otpCode, statusCode, err := u.userUsecase.SendOTP(userSendOTP.Phone)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	otpKey := fmt.Sprintf("otp:%s", userSendOTP.Phone)

	go u.redisCache.Set(otpKey, otpCode)

	slog.Info("OTP: ", otpCode)
	return NewSuccessResponse(ctx, statusCode, "OTP sent successfully", nil)
}

func (u *UsersHandler) ResendOTP(ctx echo.Context) error {
	var userSendOTP requests.UserSendOTPRequest

	err := helpers.BindAndValidate(ctx, &userSendOTP)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	var otpKey = fmt.Sprintf("otp:%s", userSendOTP.Phone)
	var attemptKey = fmt.Sprintf("attempt:%s", userSendOTP.Phone)

	attemptCount, err := u.redisCache.Get(attemptKey)
	if err != nil && err != redis.Nil {
		return NewErrorResponse(ctx, http.StatusInternalServerError, "Error checking attempt count")
	}

	if attemptCount != nil && int64(attemptCount.(float64)) > 2 {
		return NewErrorResponse(ctx, http.StatusTooManyRequests, "Maximum OTP attempts reached")
	}

	go u.redisCache.Incr(attemptKey)

	otpCode, statusCode, err := u.userUsecase.SendOTP(userSendOTP.Phone)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	go u.redisCache.Set(otpKey, otpCode)

	return NewSuccessResponse(ctx, statusCode, "OTP resent successfully", nil)
}

func (u *UsersHandler) ResetAttempts(ctx echo.Context) error {
	var userSendOTP requests.UserSendOTPRequest

	err := helpers.BindAndValidate(ctx, &userSendOTP)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	var attemptKey = fmt.Sprintf("attempt:%s", userSendOTP.Phone)

	go u.redisCache.Delete(attemptKey)

	return NewSuccessResponse(ctx, http.StatusOK, "Attempt reset successfully", nil)
}

func (u *UsersHandler) VerifyOTP(ctx echo.Context) error {
	var userVerifyOTP requests.UserVerifyOTPRequest

	err := helpers.BindAndValidate(ctx, &userVerifyOTP)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	otpKey := fmt.Sprintf("otp:%s", userVerifyOTP.Phone)
	otpCode, err := u.redisCache.Get(otpKey)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "OTP not found")
	}

	statusCode, err := u.userUsecase.VerifyOTP(userVerifyOTP.OTP, otpCode.(string))
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	go u.redisCache.Delete(otpKey)

	userExists, statusCode, err := u.userUsecase.FindByPhone(userVerifyOTP.Phone)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			verifiedKey := fmt.Sprintf("verified:%s", userVerifyOTP.Phone)
			go u.redisCache.Set(verifiedKey, true)

			return NewSuccessResponse(ctx, statusCode, "OTP verified successfully", nil)

		}
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	outDom, statusCode, err := u.userUsecase.Login(userExists)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}
	helpers.WriteCookie(ctx, "refresh_token", outDom.RefreshToken)

	return NewSuccessResponse(ctx, statusCode, "User logged in successfully", responses.FromUserDomain(outDom))

}

func (u *UsersHandler) Register(ctx echo.Context) error {
	var userRegisterRequest requests.UserRegisterRequest

	err := helpers.BindAndValidate(ctx, &userRegisterRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	verifiedKey := fmt.Sprintf("verified:%s", userRegisterRequest.Phone)
	verified, err := u.redisCache.Get(verifiedKey)
	if err != nil || !verified.(bool) {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Phone number is not verified")
	}

	userDomain := userRegisterRequest.ToDomain()

	outDomain, statusCode, err := u.userUsecase.Save(userDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	go u.redisCache.Delete(verifiedKey)

	outDom, statusCode, err := u.userUsecase.Login(outDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}
	helpers.WriteCookie(ctx, "refresh_token", outDom.RefreshToken)

	return NewSuccessResponse(ctx, statusCode, "User logged in successfully", responses.FromUserDomain(outDom))
}

func (u *UsersHandler) RefreshToken(ctx echo.Context) error {
	refreshToken, err := helpers.ReadCookie(ctx, "refresh_token")
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "refresh_token not found")
	}

	outDom, statusCode, err := u.userUsecase.RefreshToken(refreshToken)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	helpers.WriteCookie(ctx, "refresh_token", outDom.RefreshToken)

	return NewSuccessResponse(ctx, statusCode, "Token refreshed successfully", responses.FromUserDomain(outDom))
}

func (u *UsersHandler) GetMe(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	if &jwtClaims == nil {
		return NewErrorResponse(ctx, http.StatusUnauthorized, "User not found")
	}

	user, statusCode, err := u.userUsecase.FindByID(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "User found", responses.FromUserDomain(user))
}
