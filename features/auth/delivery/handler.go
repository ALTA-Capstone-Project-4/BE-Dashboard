package delivery

import (
	"warehouse/features/auth"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authUsecase auth.UsecaseInterface
}

func New(e *echo.Echo, usecase auth.UsecaseInterface) {

	handler := AuthDelivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Auth)

}

// function login
func (delivery *AuthDelivery) Auth(c echo.Context) error {
	var req AuthRequest

	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed bind data"))
	}

	str, role, name := delivery.authUsecase.LoginAuthorized(req.Email, req.Password)

	if str == "please input email and password" || str == "email not found" || str == "wrong password" {
		return c.JSON(400, helper.FailedResponseHelper(str))

	} else if str == "your account unverified" {
		return c.JSON(400, helper.FailedResponseHelper(str))

	} else if str == "error to created token" {
		return c.JSON(400, helper.FailedResponseHelper(str))

	} else {
		return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", FromCore(str, role, name)))
	}
}
