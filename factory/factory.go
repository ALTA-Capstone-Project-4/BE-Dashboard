package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authData "warehouse/features/auth/data"
	authDelivery "warehouse/features/auth/delivery"
	authUsecase "warehouse/features/auth/usecase"

	userData "warehouse/features/user/data"
	userDelivery "warehouse/features/user/delivery"
	userUsecase "warehouse/features/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authDataFactory := authData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)
}
