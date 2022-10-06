package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"warehouse/features/lahan"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type LahanDelivery struct {
	lahanUsecase lahan.UsecaseInterface
}

func New(e *echo.Echo, usecase lahan.UsecaseInterface) {
	handler := &LahanDelivery{
		lahanUsecase: usecase,
	}

	e.POST("/lahan", handler.PostLahan, middlewares.JWTMiddleware())
	e.GET("/lahan/:id", handler.GetDetailLahan, middlewares.JWTMiddleware())
	e.PUT("/lahan/:id", handler.PutLahan, middlewares.JWTMiddleware())
	e.DELETE("/lahan/:id", handler.DeleteLahan, middlewares.JWTMiddleware())
	e.GET("/penitip/lahan", handler.GetLahanClient, middlewares.JWTMiddleware())
}

func (delivery *LahanDelivery) PostLahan(c echo.Context) error {
	id, role, errToken := middlewares.ExtractToken(c)

	if role != "mitra" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}

	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	var dataLahan LahanRequest
	errBind := c.Bind(&dataLahan)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	imageData, imageInfo, imageErr := c.Request().FormFile("foto_lahan")

	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get foto_lahan"))
	}

	imageExtension, err_image_extension := helper.CheckFileExtension(imageInfo.Filename)
	if err_image_extension != nil {
		return c.JSON(400, helper.FailedResponseHelper("foto_lahan extension error"))
	}

	err_file_size := helper.CheckFileSize(imageInfo.Size)
	if err_file_size != nil {
		return c.JSON(400, helper.FailedResponseHelper("foto_lahan size error"))
	}

	imageName := strconv.Itoa(id) + time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

	image, errUploadImg := helper.UploadFileToS3("lahanimage", imageName, "images", imageData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(400, helper.FailedResponseHelper("failed to upload foto_lahan"))
	}

	dataLahan.FotoLahan = image

	row_postLahan, err_postLahan := delivery.lahanUsecase.PostLahan(toCore(dataLahan), id)

	if row_postLahan != 1 || err_postLahan != nil {
		return c.JSON(500, helper.FailedResponseHelper("error add data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success add data"))
}

func (delivery *LahanDelivery) GetDetailLahan(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role == "admin" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}
	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	id := c.Param("id")
	idlahan, _ := strconv.Atoi(id)

	data, err := delivery.lahanUsecase.GetDetailLahan(idlahan, role)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}
	if data.ID == 0 {
		return c.JSON(400, helper.FailedResponseHelper("there is no data mitra"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))
}

func (delivery *LahanDelivery) PutLahan(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "mitra" {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}
	if errToken != nil {
		return c.JSON(400, helper.FailedResponseHelper("Invalid token"))
	}

	idlahan := c.Param("id")
	idCnv, _ := strconv.Atoi(idlahan)

	var dataUpdate LahanRequest

	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}
	eventCore := toCore(dataUpdate)

	fotoData, fotoInfo, fotoErr := c.Request().FormFile("foto_lahan")

	if fotoData != nil {
		if fotoErr == http.ErrMissingFile || fotoErr != nil {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get photo profile"))
		}

		fotoExtension, err_foto_extension := helper.CheckFileExtension(fotoInfo.Filename)
		if err_foto_extension != nil {
			return c.JSON(400, helper.FailedResponseHelper("photo profile extension error"))
		}

		err_foto_size := helper.CheckFileSize(fotoInfo.Size)
		if err_foto_size != nil {
			return c.JSON(400, helper.FailedResponseHelper("photo profile size error"))
		}

		fotoName := strconv.Itoa(token) + time.Now().Format("2006-01-02 15:04:05") + "." + fotoExtension

		foto, errUploadFoto := helper.UploadFileToS3("fotoprofileimage", fotoName, "images", fotoData)

		if errUploadFoto != nil {
			fmt.Println(errUploadFoto)
			return c.JSON(400, helper.FailedResponseHelper("failed to upload foto lahan"))
		}
		eventCore.FotoLahan = foto
	}

	row, err := delivery.lahanUsecase.PutLahan(idCnv, token, eventCore)
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("Unauthorized"))
	}
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success update data"))
}

func (delivery *LahanDelivery) DeleteLahan(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}
	if role != "mitra" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unauthorized"))
	}

	id := c.Param("id")
	idCnv, _ := strconv.Atoi(id)

	var data lahan.Core
	row, err := delivery.lahanUsecase.DeleteLahan(idCnv, token, data)
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("Unauthorized"))
	}
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed delete"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete"))
}

func (delivery *LahanDelivery) GetLahanClient(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	data, err := delivery.lahanUsecase.GetLahanClient(token)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromLahanPenitipList(data)))
}
