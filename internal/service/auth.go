package service

import (
	"errors"
	"time"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/dto"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/storage"
	"voice-out-be/internal/util"
	"voice-out-be/internal/voerrors"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type authService struct {
	config  Config
	storage storage.Storage
}

type JwtClaims struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func NewAuthService(config Config, storage storage.Storage) *authService {

	return &authService{
		storage: storage,
		config:  config,
	}
}

func (a *authService) RegisterUser(registerUserReq *request.RegisterUserRequest) (code response.Code, token string, err error) {

	a.config.Logger().Info("RegisterUser: finding user by email")
	existingUser, err := a.storage.FindUserByEmail(registerUserReq.Email)

	if err != nil {
		if !errors.Is(err, voerrors.ErrNotFound) {
			a.config.Logger().Error("RegisterUser: error finding user by email", zap.Error(err))
			return response.ServerError, "", err
		}
	}

	if existingUser != nil {
		a.config.Logger().Error("RegisterUser: user already exists", zap.Error(err))
		return response.DuplicateUser, "", voerrors.ErrDuplicateUser
	}

	hashedPassword, err := util.HashPassword(registerUserReq.Password)

	if err != nil {
		a.config.Logger().Error("RegisterUser: unable to hash password", zap.Error(err))
		return response.ServerError, "", err
	}

	userDto := &dto.RegisterUser{
		FirstName: registerUserReq.FirstName,
		LastName:  registerUserReq.LastName,
		Username:  registerUserReq.Username,
		Email:     registerUserReq.Email,
		Password:  hashedPassword,
	}

	user, err := a.storage.CreateUser(userDto)

	if err != nil {
		a.config.Logger().Error("RegisterUser: unable to save user to db", zap.Error(err))
		return response.ServerError, "", err
	}

	token, err = a.GenerateJWT(user.Username, user.Email, user.ID)

	if err != nil {
		return response.ServerError, "", err
	}

	return response.Success, token, nil
}

func (a *authService) Login(loginReq *request.LoginRequest) (code response.Code, token string, err error) {

	a.config.Logger().Info("Login:  finding user by email")
	existingUser, err := a.storage.FindUserByEmail(loginReq.Email)

	if err != nil {
		if !errors.Is(err, voerrors.ErrNotFound) {
			a.config.Logger().Error("Login: error finding user by email", zap.Error(err))
			return response.ServerError, "", err
		}
	}

	if existingUser == nil {
		a.config.Logger().Error("Login: email does not exist", zap.Error(err))
		return response.NotFound, "", voerrors.ErrNotFound
	}

	a.config.Logger().Info("Login: checking password hash")
	isPasswordMatch := util.CheckPasswordHash(loginReq.Password, existingUser.Password)

	if !isPasswordMatch {
		a.config.Logger().Error("Login: incorrect password", zap.Error(err))
		return response.BadRequest, "", voerrors.ErrBadRequest
	}

	token, err = a.GenerateJWT(existingUser.Username, existingUser.Email, existingUser.ID)

	if err != nil {
		return response.ServerError, "", err
	}

	return response.Success, token, nil
}

func (a *authService) GenerateJWT(name string, email string, userID string) (jwtToken string, err error) {
	a.config.Logger().Info("GenerateJWT: generating jwt token")

	claims := JwtClaims{
		Name:   name,
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    constants.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(constants.JWT_EXPIRATION_DURATION).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		a.config.Logger().Error("GenerateJWT: error generating jwt token", zap.Error(err))
		return "", err
	}

	return t, nil

}
