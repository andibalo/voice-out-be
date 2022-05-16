package service

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"time"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/dto"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/storage"
	"voice-out-be/internal/util"
	"voice-out-be/internal/voerrors"
)

type authService struct {
	config  Config
	storage storage.Storage
}

type jwtClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewAuthService(config Config, storage storage.Storage) *authService {

	return &authService{
		storage: storage,
		config:  config,
	}
}

func (a *authService) RegisterUser(registerUserReq *request.RegisterUserRequest) (response.Code, error) {

	a.config.Logger().Info("RegisterUser: finding user by email")
	existingUser, err := a.storage.FindUserByEmail(registerUserReq.Email)

	if err != nil {
		a.config.Logger().Error("RegisterUser: error finding user by email", zap.Error(err))
		return response.ServerError, err
	}

	if existingUser != nil {
		a.config.Logger().Error("RegisterUser: user already exists", zap.Error(err))
		return response.DuplicateUser, voerrors.ErrDuplicateUser
	}

	hashedPassword, err := util.HashPassword(registerUserReq.Password)

	if err != nil {
		a.config.Logger().Error("RegisterUser: unable to hash password", zap.Error(err))
		return response.ServerError, err
	}

	user := &dto.RegisterUser{
		FirstName: registerUserReq.FirstName,
		LastName:  registerUserReq.LastName,
		Username:  registerUserReq.Username,
		Email:     registerUserReq.Email,
		Password:  hashedPassword,
	}

	_, err = a.storage.CreateUser(user)

	if err != nil {
		a.config.Logger().Error("RegisterUser: unable to save user to db", zap.Error(err))
		return response.ServerError, err
	}

	return response.Success, nil
}

func (a *authService) GenerateJWT(name string, email string) (jwtToken string, err error) {
	claims := jwtClaims{
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    constants.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(constants.JWT_EXPIRATION_DURATION).Unix(),
		},
	}

	a.config.Logger().Info("GenerateJWT: generating jwt token", zap.Error(err))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("TEST"))
	if err != nil {
		a.config.Logger().Error("GenerateJWT: error generating jwt token", zap.Error(err))
		return "", err
	}

	return t, nil

}
