package services

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/mframadann/gourl/domain/auth/dto"
	"github.com/mframadann/gourl/domain/auth/repositories"
	"github.com/mframadann/gourl/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(user dto.Register) utils.Response
	SignIn(payload dto.SignIn) utils.Response
}

type authService struct {
	authRepo repositories.AuthRepository
}

func (service *authService) Register(user dto.Register) utils.Response {
	var response utils.Response
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		response.Message = "Oops! something when wrong."
		response.StatusCode = 500
		return response
	}

	user.Password = string(hashedPass)
	err = service.authRepo.Register(user)

	if err != nil {
		response.Message = "Failed Register."
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Message = "Register successfully"
	response.StatusCode = http.StatusOK
	return response
}

func (service *authService) SignIn(payload dto.SignIn) utils.Response {
	var res utils.Response
	data, err := service.authRepo.SignIn(payload)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		res.Message = "Record not found."
		res.StatusCode = http.StatusNotFound
		return res
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		res.Message = "Login failed"
		res.StatusCode = http.StatusBadRequest
		return res
	}

	accessToken, err := utils.CreateToken(&utils.TokenPayload{
		Sub:       data.EmailAddress,
		Exp:       time.Now().Add(time.Hour).Unix(),
		SecretKey: []byte(os.Getenv("ACCESS_SECRET_KEY")),
	})

	refreshToken, errRefreshToken := utils.CreateToken(&utils.TokenPayload{
		Sub:       data.EmailAddress,
		Exp:       time.Now().Add(time.Hour * 24).Unix(),
		SecretKey: []byte(os.Getenv("REFRESH_SECRET_KEY")),
	})

	if err != nil && errRefreshToken != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Error creating token" + err.Error()
		return res
	}

	res.Message = "Login successfully."
	res.StatusCode = http.StatusOK
	res.Data = map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return res
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{authRepo: repositories.NewAuthRepository(db)}
}
