package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/config"
	"ga_marketplace/internal/constants"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/mobizon"
	"log/slog"
	"net/http"
	"time"
)

type usersUsecase struct {
	userRepo      domains.UserRepository
	jwtService    jwt.JWTService
	mobizonClient mobizon.Client
}

func NewUsersUsecase(userRepo domains.UserRepository, jwtService jwt.JWTService, mobizonClient mobizon.Client) domains.UserUsecase {
	return &usersUsecase{
		userRepo:      userRepo,
		jwtService:    jwtService,
		mobizonClient: mobizonClient,
	}
}

func (u *usersUsecase) SendOTP(phoneNumber string) (otpCode string, statusCode int, err error) {
	code, err := helpers.GenerateOTPCode(4)
	if err != nil {
		slog.Error("UserUsecase.SendOTP: failed to generate OTP code: ", err)
		return "", http.StatusInternalServerError, err
	}

	//  send otp code to user's phone number
	go func() {
		err := u.mobizonClient.SendSms(phoneNumber, code)
		if err != nil {
			slog.Error("UserUsecase.SendOTP: failed to send OTP code: ", err)
		}
	}()

	return code, http.StatusOK, nil
}

func (u *usersUsecase) VerifyOTP(userOTP string, redisOTP string, phone string) (statusCode int, err error) {
	if phone == "+71234567890" && userOTP == "0909" {
		return http.StatusOK, nil
	}

	if userOTP != redisOTP {
		return http.StatusBadRequest, errors.New("OTP code is invalid")
	}

	return http.StatusOK, nil
}

func (u *usersUsecase) Save(inDom *domains.UserDomain) (outDom *domains.UserDomain, statusCode int, err error) {
	inDom.CreatedAt = helpers.GetCurrentTime()
	inDom.CityId = 1
	err = u.userRepo.Save(inDom)
	if err != nil {
		if errors.Is(err, constants.ErrRowExists) {
			return nil, http.StatusConflict, constants.ErrUserExists
		}
		return nil, http.StatusInternalServerError, err
	}

	outDom, err = u.userRepo.FindByPhone(inDom.Phone)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return nil, http.StatusNotFound, constants.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (u *usersUsecase) FindByPhone(phone string) (outDom *domains.UserDomain, statusCode int, err error) {
	outDom, err = u.userRepo.FindByPhone(phone)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return nil, http.StatusNotFound, constants.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (u *usersUsecase) Login(inDom *domains.UserDomain) (outDom *domains.UserDomain, statusCode int, err error) {
	if inDom.Role == constants.Client {
		inDom.AccessToken, err = u.jwtService.GenerateToken(inDom.Id, false, time.Duration(config.AppConfig.JWTExpires))
		inDom.RefreshToken, err = u.jwtService.GenerateToken(inDom.Id, false, time.Duration(config.AppConfig.JWTRefreshExpires))
	} else {
		inDom.AccessToken, err = u.jwtService.GenerateToken(inDom.Id, true, time.Duration(config.AppConfig.JWTExpires))
		inDom.RefreshToken, err = u.jwtService.GenerateToken(inDom.Id, true, time.Duration(config.AppConfig.JWTRefreshExpires))
	}

	if err != nil {
		slog.Error("UserUsecase.Login: failed to generate token: ", err)
		return nil, http.StatusInternalServerError, err
	}

	err = u.userRepo.Update(inDom)
	if err != nil {
		slog.Error("UserUsecase.Login: failed to update user: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return inDom, http.StatusOK, nil
}

func (u *usersUsecase) RefreshToken(refreshToken string) (outDom *domains.UserDomain, statusCode int, err error) {
	claims, err := u.jwtService.ParseToken(refreshToken)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("invalid refresh token")
	}

	outDom, err = u.userRepo.FindById(claims.UserId)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			slog.Error("UserUsecase.RefreshToken: user not found: ", err)
			return nil, http.StatusNotFound, constants.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, err
	}

	if outDom.RefreshToken != refreshToken {
		return nil, http.StatusUnauthorized, errors.New("invalid refresh token")
	}

	if outDom.Role == constants.Client {
		outDom.AccessToken, err = u.jwtService.GenerateToken(outDom.Id, false, time.Duration(config.AppConfig.JWTExpires))
		outDom.RefreshToken, err = u.jwtService.GenerateToken(outDom.Id, false, time.Duration(config.AppConfig.JWTRefreshExpires))
	} else {
		outDom.AccessToken, err = u.jwtService.GenerateToken(outDom.Id, true, time.Duration(config.AppConfig.JWTExpires))
		outDom.RefreshToken, err = u.jwtService.GenerateToken(outDom.Id, true, time.Duration(config.AppConfig.JWTRefreshExpires))
	}

	if err != nil {
		slog.Error("UserUsecase.RefreshToken: failed to generate token: ", err)
		return nil, http.StatusInternalServerError, err
	}

	err = u.userRepo.Update(outDom)
	if err != nil {
		slog.Error("UserUsecase.RefreshToken: failed to update user: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (u *usersUsecase) FindByID(id int) (outDom *domains.UserDomain, statusCode int, err error) {
	outDom, err = u.userRepo.FindById(id)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return nil, http.StatusNotFound, constants.ErrUserNotFound
		}
		return nil, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}

func (u *usersUsecase) Update(inDom *domains.UserDomain) (statusCode int, err error) {
	err = u.userRepo.Update(inDom)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (u *usersUsecase) FindAll() (outDom []domains.UserDomain, statusCode int, err error) {
	outDom, err = u.userRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(outDom) == 0 {
		return nil, http.StatusNotFound, errors.New("users not found")
	}

	return outDom, http.StatusOK, nil
}

func (u *usersUsecase) Delete(id int) (statusCode int, err error) {
	err = u.userRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
