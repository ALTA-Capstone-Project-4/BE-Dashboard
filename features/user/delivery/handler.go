package delivery

import (
	"warehouse/features/user"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

// func New(e *echo.Echo, usecase user.UsecaseInterface) {
// 	handler := &UserDelivery{
// 		userUsecase: usecase,
// 	}

// }
