package delivery

import (
	"fmt"
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

	e.POST("/register", handler.PostUser)
	e.GET("/mitra/:id", handler.GetMitraByAdmin, middlewares.JWTMiddleware())
	e.GET("/mitra", handler.GetMitra, middlewares.JWTMiddleware())
	e.PUT("/mitra", handler.PutMitra, middlewares.JWTMiddleware())
	// e.DELETE("/users/:id", handler.DeleteUser, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) PostUser(c echo.Context) error {
	var userRegister UserRequest

	errBind := c.Bind(&userRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}
	fmt.Println(userRegister)

	// f, err := c.FormFile("file_ktp")
	// if err != nil {
	// 	return c.JSON(400, helper.FailedResponseHelper("error bind ktp file"))
	// }

	// blobFile, err := f.Open()
	// if err != nil {
	// 	return c.JSON(400, helper.FailedResponseHelper("error open ktp file"))
	// }

	// err = config.Uploader.UploadFile(blobFile, f.Filename)
	// if err != nil {

	// 	return c.JSON(500, helper.FailedResponseHelper(err.Error()))
	// }

	row, err := delivery.userUsecase.PostUser(toCore(userRegister))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *UserDelivery) GetMitraByAdmin(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	id := c.Param("id")
	idCnv, _ := strconv.Atoi(id)

	data, err := delivery.userUsecase.GetMitra(idCnv)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))
}

func (delivery *UserDelivery) GetMitra(c echo.Context) error {
	token, _, errToken := middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	data, err := delivery.userUsecase.GetMitra(token)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))
}

func (delivery *UserDelivery) PutMitra(c echo.Context) error {
	token, _, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	var dataUpdate UserRequest

	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row, err := delivery.userUsecase.PutMitra(token, toCore(dataUpdate))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success update data"))
}

// func (delivery *UserDelivery) DeleteUser(c echo.Context) error {
// 	_, role, errToken := middlewares.ExtractToken(c)

// 	if role == "mitra" {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
// 	}
// 	if errToken != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
// 	}

// 	id := c.Param("id")
// 	idCnv, errId := strconv.Atoi(id)
// 	if errId != nil {
// 		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
// 	}

// 	query := c.QueryParam("admin")
// 	query1 := c.QueryParam("client")

// 	row, err := delivery.userUsecase.DeleteUser(idCnv, query, query1)
// 	if err != nil || row != 1 {
// 		return c.JSON(400, helper.FailedResponseHelper("failed delete"))
// 	}
// 	return c.JSON(200, helper.SuccessResponseHelper("success delete"))

// }
