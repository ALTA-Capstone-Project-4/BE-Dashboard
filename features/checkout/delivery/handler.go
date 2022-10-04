package delivery

import (
	"fmt"
	"net/http"
	"time"
	"warehouse/features/checkout"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type CheckoutDelivery struct {
	checkoutUsecase checkout.UsecaseInterface
}

func New(e *echo.Echo, usecase checkout.UsecaseInterface) {
	handler := &CheckoutDelivery{
		checkoutUsecase: usecase,
	}
	e.POST("/order", handler.PostCheckoutByFav, middlewares.JWTMiddleware())
}

func (delivery *CheckoutDelivery) PostCheckoutByFav(c echo.Context) error {
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

	fmt.Println(dataCheckout.MulaiSewa)
	dataCore := toCore(dataCheckout)
	dataCore.MulaiSewa = mulaiSewa
	dataCore.AkhirSewa = akhirSewa
	dataCore.UserID = token
	fmt.Println(dataCore.MulaiSewa, "dataCore mulai sewa")

	row, err := delivery.checkoutUsecase.PostCheckoutByFav(dataCore)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}
