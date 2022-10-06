package delivery

import (
	"strconv"
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

	e.GET("/gudang", handler.GetAllGudang, middlewares.JWTMiddleware())
	e.POST("/gudang", handler.PostGudang, middlewares.JWTMiddleware())
	e.GET("/gudang/:id/lahan", handler.GetGudangByID, middlewares.JWTMiddleware())

}

func (delivery *GudangDelivery) GetAllGudang(c echo.Context) error {

	_, role, errToken := middlewares.ExtractToken(c)

	if role == "mitra" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}
	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	page := c.QueryParam("page")
	pageCnv, errPage := strconv.Atoi(page)

	if errPage != nil && page != "" {
		return c.JSON(400, helper.FailedResponseHelper("page param must be number"))
	}

	data, err := delivery.gudangUsecase.GetAllGudang(pageCnv) //untuk homepage
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

	if row_postGudang != 1 || err_postGudang != nil {
		return c.JSON(500, helper.FailedResponseHelper("error add data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success add data"))
}

func (delivery *GudangDelivery) GetGudangByID(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role == "admin" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}
	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	gudang_id := c.Param("id")
	idCnv, _ := strconv.Atoi(gudang_id)
	data, err := delivery.gudangUsecase.GetGudangByID(idCnv)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))
}
