package delivery

import (
	"warehouse/features/lahan"

	"github.com/labstack/echo/v4"
)

type LahanDelivery struct {
	lahanUsecase lahan.UsecaseInterface
}

func New(e *echo.Echo, usecase lahan.UsecaseInterface) {
	// handler := &LahanDelivery{
	// 	lahanUsecase: usecase,
	// }

}
