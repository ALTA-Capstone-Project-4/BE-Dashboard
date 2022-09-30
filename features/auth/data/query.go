package data

import (
	"warehouse/features/auth"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.DataInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) LoginUser(email string) (auth.Core, error) {
	var data User

	if data.Role == "mitra" {
		txMitra := repo.db.Where("email = ? AND status = ?", email, "verified").First(&data)
		if txMitra.Error != nil {
			return auth.Core{}, txMitra.Error
		}
		if txMitra.RowsAffected != 1 {
			return auth.Core{}, txMitra.Error
		}

		return toCore(data), nil

	} else {

		txClient := repo.db.Where("email = ?", email).First(&data)
		if txClient.Error != nil {
			return auth.Core{}, txClient.Error
		}
		if txClient.RowsAffected != 1 {
			return auth.Core{}, txClient.Error
		}

		var dataUser = toCore(data)
		return dataUser, nil
	}
}
