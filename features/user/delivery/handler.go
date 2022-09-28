package delivery

import (
	"net/http"
	"strconv"
	"warehouse/features/user"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.POST("/register", handler.PostUser, middlewares.JWTMiddleware())
	e.GET("/admin/mitra/:id", handler.GetMitraId, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) PostUser(c echo.Context) error {
	// _, role, errToken := middlewares.ExtractToken(c)

	// if role != "admin" {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	// }
	// if errToken != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	// }

	var userRegister UserRequest

	errBind := c.Bind(&userRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row, err := delivery.userUsecase.PostUser(toCore(userRegister))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *UserDelivery) GetMitraId(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role == "client" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	id := c.Param("id")
	idCnv, _ := strconv.Atoi(id)

	data, err := delivery.userUsecase.GetMitraId(idCnv)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))
}
