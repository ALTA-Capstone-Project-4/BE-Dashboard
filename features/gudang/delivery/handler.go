package delivery

import (
	"warehouse/features/gudang"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type GudangDelivery struct {
	gudangUsecase gudang.UsecaseInterface
}

func New(e *echo.Echo, usecase gudang.UsecaseInterface) {
	handler := &GudangDelivery{
		gudangUsecase: usecase,
	}

	e.PUT("/gudang", handler.PutGudang, middlewares.JWTMiddleware())
	e.GET("/gudang", handler.GetAllGudang)
	e.POST("/gudang", handler.PostGudang, middlewares.JWTMiddleware())

}

func (delivery *GudangDelivery) PutGudang(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "mitra" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}
	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	var dataUpdate GudangRequest

	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row, err := delivery.gudangUsecase.PutGudang(token, toCore(dataUpdate))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success update data"))
}

func (delivery *GudangDelivery) GetAllGudang(c echo.Context) error {
	data, err := delivery.gudangUsecase.GetAllGudang() //untuk homepage
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore_toHomeList(data)))
}

func (delivery *GudangDelivery) PostGudang(c echo.Context) error {
	id, role, errToken := middlewares.ExtractToken(c)

	if role != "mitra" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}
	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	var dataGudang GudangRequest
	errBind := c.Bind(&dataGudang)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	dataGudang.UserID = id

	row_postGudang, err_postGudang := delivery.gudangUsecase.PostGudang(toCore(dataGudang))
	if err_postGudang.Error() == "your account unverified" {
		return c.JSON(500, helper.FailedResponseHelper(err_postGudang.Error()))
	}
	if row_postGudang != 1 || err_postGudang != nil {
		return c.JSON(500, helper.FailedResponseHelper("error add data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success add data"))
}
