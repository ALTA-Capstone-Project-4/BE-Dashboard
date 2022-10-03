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

	gudangData "warehouse/features/gudang/data"
	gudangDelivery "warehouse/features/gudang/delivery"
	gudangUsecase "warehouse/features/gudang/usecase"

	lahanData "warehouse/features/lahan/data"
	lahanDelivery "warehouse/features/lahan/delivery"
	lahanUsecase "warehouse/features/lahan/usecase"

	favoriteData "warehouse/features/favorite/data"
	favoriteDelivery "warehouse/features/favorite/delivery"
	favoriteUsecase "warehouse/features/favorite/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authDataFactory := authData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	gudangDataFactory := gudangData.New(db)
	gudangUsecaseFactory := gudangUsecase.New(gudangDataFactory)
	gudangDelivery.New(e, gudangUsecaseFactory)

	lahanDataFactory := lahanData.New(db)
	lahanUsecaseFactory := lahanUsecase.New(lahanDataFactory)
	lahanDelivery.New(e, lahanUsecaseFactory)

	favoriteDataFactory := favoriteData.New(db)
	favoriteUsecaseFactory := favoriteUsecase.New(favoriteDataFactory)
	favoriteDelivery.New(e, favoriteUsecaseFactory)
}
