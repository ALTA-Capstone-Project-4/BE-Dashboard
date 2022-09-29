package migration

import (
	gudangModel "warehouse/features/gudang/data"
	userModel "warehouse/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&gudangModel.Gudang{})
}
