package migration

import (
	checkoutModel "warehouse/features/checkout/data"
	favoriteModel "warehouse/features/favorite/data"
	gudangModel "warehouse/features/gudang/data"
	lahanModel "warehouse/features/lahan/data"
	userModel "warehouse/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&gudangModel.Gudang{})
	db.AutoMigrate(&lahanModel.Lahan{})
	db.AutoMigrate(&favoriteModel.Favorite{})
	db.AutoMigrate(&checkoutModel.Checkout{})
}
