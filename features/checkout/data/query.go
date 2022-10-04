package data

import (
	"warehouse/features/checkout"

	"gorm.io/gorm"
)

type checkoutData struct {
	db *gorm.DB
}

func New(db *gorm.DB) checkout.DataInterface {
	return &checkoutData{
		db: db,
	}
}

func (repo *checkoutData) AddCheckoutByFav(data checkout.Core) (int, error) {

	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
