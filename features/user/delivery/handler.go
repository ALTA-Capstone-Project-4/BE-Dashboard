package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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
	e.GET("/mitra/unverify", handler.GetMitraUnverif, middlewares.JWTMiddleware())
	e.PUT("/mitra/verify/:id", handler.PutVerify, middlewares.JWTMiddleware())
	e.GET("/mitra/verified", handler.GetMitraVerified, middlewares.JWTMiddleware())
	e.GET("/mitra/:id", handler.GetMitraByAdmin, middlewares.JWTMiddleware())
	e.GET("/mitra", handler.GetMitra, middlewares.JWTMiddleware())
	e.PUT("/mitra", handler.PutMitra, middlewares.JWTMiddleware())
	e.DELETE("/mitra/:id", handler.DeleteMitra, middlewares.JWTMiddleware())
	e.GET("/penitip", handler.GetClient, middlewares.JWTMiddleware())
	e.PUT("/penitip", handler.PutClient, middlewares.JWTMiddleware())
	e.DELETE("/penitip", handler.DeleteClient, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) PostUser(c echo.Context) error {
	var userRegister UserRequest

	errBind := c.Bind(&userRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}
	fmt.Println(userRegister)

	if userRegister.Role == "mitra" && userRegister.Status == "" {
		userRegister.Status = "unverified"
	}

	if userRegister.Role == "penitip" || userRegister.Role == "admin" {
		userRegister.FileKTP = ""

	} else {

		imageData, imageInfo, imageErr := c.Request().FormFile("file_ktp")

		if imageErr == http.ErrMissingFile || imageErr != nil {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get file_ktp"))
		}

		imageExtension, err_image_extension := helper.CheckFileExtension(imageInfo.Filename)
		if err_image_extension != nil {
			return c.JSON(400, helper.FailedResponseHelper("file_ktp extension error"))
		}

		err_file_size := helper.CheckFileSize(imageInfo.Size)
		if err_file_size != nil {
			return c.JSON(400, helper.FailedResponseHelper("file_ktp size error"))
		}

		imageName := time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

		image, errUploadImg := helper.UploadFileToS3("ktpimage", imageName, "images", imageData)

		if errUploadImg != nil {
			fmt.Println(errUploadImg)
			return c.JSON(400, helper.FailedResponseHelper("failed to upload file_ktp"))
		}

		eventCore := toCore(userRegister)
		eventCore.FileKTP = image

		row, err := delivery.userUsecase.PostUser(eventCore)
		if err != nil {
			return c.JSON(500, helper.FailedResponseHelper("error insert data"))
		}
		if row != 1 {
			return c.JSON(500, helper.FailedResponseHelper("error insert data"))
		}

		return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
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

func (delivery *UserDelivery) GetMitraUnverif(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	data, err := delivery.userUsecase.GetMitraUnverif()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCoreList(data)))
}

func (delivery *UserDelivery) PutVerify(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	id := c.Param("id")
	idCnv, _ := strconv.Atoi(id)

	var verify UserRequest
	errBind := c.Bind(&verify)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row, err := delivery.userUsecase.PutVerify(idCnv, toCore(verify))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update status"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update status"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success update status"))

}

func (delivery *UserDelivery) GetMitraVerified(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	data, err := delivery.userUsecase.GetMitraVerified()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCoreList(data)))
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
	if data.ID == 0 {
		return c.JSON(400, helper.FailedResponseHelper("there is no data mitra"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))
}

func (delivery *UserDelivery) GetMitra(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}
	if role == "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unauthorized"))
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

	fotoData, fotoInfo, fotoErr := c.Request().FormFile("photo")

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

	id, _, _ := middlewares.ExtractToken(c)
	fotoName := strconv.Itoa(id) + time.Now().Format("2006-01-02 15:04:05") + "." + fotoExtension

	foto, errUploadFoto := helper.UploadFileToS3("fotoprofileimage", fotoName, "images", fotoData)

	if errUploadFoto != nil {
		fmt.Println(errUploadFoto)
		return c.JSON(400, helper.FailedResponseHelper("failed to upload photo profile"))
	}

	eventCore := toCore(dataUpdate)
	eventCore.Photo = foto

	row, err := delivery.userUsecase.PutMitra(token, eventCore)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success update data"))
}

func (delivery *UserDelivery) DeleteMitra(c echo.Context) error {
	_, role, errToken := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	id := c.Param("id")
	idCnv, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, err := delivery.userUsecase.DeleteMitra(idCnv)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed delete"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete"))
}

func (delivery *UserDelivery) GetClient(c echo.Context) error {
	id, role, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}
	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unauthorized"))
	}

	data, err := delivery.userUsecase.GetClient(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCore(data)))

}

func (delivery *UserDelivery) PutClient(c echo.Context) error {
	token, _, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	var dataUpdate UserRequest

	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	var foto string
	eventCore := toCore(dataUpdate)

	if dataUpdate.Photo != "" {
		fotoData, fotoInfo, fotoErr := c.Request().FormFile("photo")

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

		id, _, _ := middlewares.ExtractToken(c)
		fotoName := strconv.Itoa(id) + time.Now().Format("2006-01-02 15:04:05") + "." + fotoExtension

		fileFoto, errUploadFoto := helper.UploadFileToS3("fotoprofileimage", fotoName, "images", fotoData)

		foto = fileFoto

		if errUploadFoto != nil {
			fmt.Println(errUploadFoto)
			return c.JSON(400, helper.FailedResponseHelper("failed to upload photo profile"))
		}
		eventCore.Photo = foto
	}

	row, err := delivery.userUsecase.PutClient(token, eventCore)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success update data"))
}

func (delivery *UserDelivery) DeleteClient(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}
	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unauthorized"))
	}

	row, err := delivery.userUsecase.DeleteClient(token)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed delete"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete"))
}
