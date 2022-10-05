package delivery

import (
	"fmt"
	"net/http"
	"time"
	"warehouse/config"
	"warehouse/features/checkout"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type CheckoutDelivery struct {
	checkoutUsecase checkout.UsecaseInterface
}

var event coreapi.Client

func New(e *echo.Echo, usecase checkout.UsecaseInterface) {

	handler := CheckoutDelivery{
		checkoutUsecase: usecase,
	}
	e.POST("/order", handler.PostCheckoutByFav, middlewares.JWTMiddleware())
	e.POST("/callback", handler.MidtransWebHook)
}

func (delivery CheckoutDelivery) PostCheckoutByFav(c echo.Context) error {
	midtrans.ServerKey = config.MidtransServerKey()
	event.New(midtrans.ServerKey, midtrans.Sandbox)

	token, role, errToken := middlewares.ExtractToken(c)

	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	var dataCheckout CheckoutRequest

	errBind := c.Bind(&dataCheckout)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	layout_date := "2006-01-02T15:04"
	mulaiSewa, err := time.Parse(layout_date, fmt.Sprintf("%sT00:00", dataCheckout.MulaiSewa))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed format tanggal mulai sewa"))
	}
	akhirSewa, err := time.Parse(layout_date, fmt.Sprintf("%sT00:00", dataCheckout.AkhirSewa))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed format tanggal akhir sewa"))
	}

	imageData, imageInfo, imageErr := c.Request().FormFile("foto")

	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get foto_barang"))
	}

	imageExtension, err_image_extension := helper.CheckFileExtension(imageInfo.Filename)
	if err_image_extension != nil {
		return c.JSON(400, helper.FailedResponseHelper("foto_barang extension error"))
	}

	err_file_size := helper.CheckFileSize(imageInfo.Size)
	if err_file_size != nil {
		return c.JSON(400, helper.FailedResponseHelper("foto_barang size error"))
	}

	imageName := time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

	image, errUploadImg := helper.UploadFileToS3("barangimage", imageName, "images", imageData)

	if errUploadImg != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to upload foto_barang"))
	}

	dataCore := toCore(dataCheckout)
	dataCore.MulaiSewa = mulaiSewa
	dataCore.AkhirSewa = akhirSewa
	dataCore.UserID = token
	dataCore.FotoBarang = image

	hargaLahan, mitra_id, errHarga := delivery.checkoutUsecase.GetDataLahan(dataCheckout.LahanID, role)

	if errHarga != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get harga lahan"))
	}

	dataCore.TotalHarga = hargaLahan * dataCore.Periode

	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	timer := currentTime.Format("15:04:05")

	orderIDPay := fmt.Sprintf("Order-%d-%s-%s", token, date, timer)
	dataCore.OrderID = orderIDPay

	inputPay := ToCoreMidtrans(dataCore)

	if dataCheckout.MetodePembayaran == "BCA" {
		dataCore.MetodePembayaran = "BCA"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	}

	if dataCheckout.MetodePembayaran == "BRI" {
		dataCore.MetodePembayaran = "BRI"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBri,
		}
	}

	if dataCheckout.MetodePembayaran == "BNI" {
		dataCore.MetodePembayaran = "BNI"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
		}
	}

	detailPay, errPay := delivery.checkoutUsecase.CreatePaymentBankTransfer(dataCheckout.LahanID, mitra_id, inputPay)

	if errPay != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data to midtrans"))
	}

	result := FromMidtransToPayment(detailPay)

	layout := "2006-01-02 15:04:05"
	trTime, _ := time.Parse(layout, detailPay.TransactionTime)
	result.TransactionTime = trTime
	result.TransactionExpire = trTime.Add(time.Hour * 24)

	dataCore.OrderID = result.OrderID
	dataCore.Status = result.TransactionStatus
	dataCore.TransactionID = result.TransactionID
	dataCore.BillNumber = result.BillNumber
	dataCore.TransactionExpire = result.TransactionExpire.String()

	row, err := delivery.checkoutUsecase.PostCheckoutByFav(dataCore)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper("success insert data", result))
}

func (delivery CheckoutDelivery) MidtransWebHook(c echo.Context) error {
	midtransRequest := MidtransHookRequest{}
	err_bind := c.Bind(&midtransRequest)
	if err_bind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	errUpdateStatusPay := delivery.checkoutUsecase.PaymentWebHook(midtransRequest.OrderID, midtransRequest.TransactionStatus)
	if errUpdateStatusPay != nil {
		return c.JSON(500, helper.FailedResponseHelper("failed to update status payment"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success to update status payment"))
}
