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
