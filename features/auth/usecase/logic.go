package usecase

import (
	"warehouse/features/auth"
	"warehouse/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authData auth.DataInterface
}

func New(data auth.DataInterface) auth.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) LoginAuthorized(email, password string) (string, string) {
	if email == "" || password == "" {
		return "please input email and password", ""
	}

	result, errEmail := usecase.authData.LoginUser(email)
	if errEmail != nil {
		return "email not found", ""
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if errPw != nil {
		return "wrong password", ""
	}

	token, errToken := middlewares.CreateToken(int(result.ID), result.Role)

	if errToken != nil {
		return "error to created token", ""
	}

	if result.Status != "verified" {
		return "your account unverified", ""
	}

	return token, result.Role

}
