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
	txEmail := repo.db.Where("email = ?", email).First(&data)
	if txEmail.Error != nil {
		return auth.Core{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		return auth.Core{}, txEmail.Error
	}

	var dataUser = toCore(data)

	return dataUser, nil
}
