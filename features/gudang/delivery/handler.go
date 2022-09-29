package delivery

import (
	"warehouse/features/gudang"
)

type GudangDelivery struct {
	gudangUsecase gudang.UsecaseInterface
}

// func New(e *echo.Echo, usecase gudang.UsecaseInterface) {
// 	handler := &GudangDelivery{
// 		gudangUsecase: usecase,
// 	}

// }

// func (delivery *GudangDelivery) PutGudang(c echo.Context) error {

// }
