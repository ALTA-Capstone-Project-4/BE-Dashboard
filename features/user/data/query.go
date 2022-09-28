package data

import (
	"warehouse/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) AddUser(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) SelectUserProfile(id int, userId int) (user.Core, error) {
	var data User

	if data.Role == "admin" || data.Role == "mitra" {
		tx := repo.db.Where("id = ?", id).Preload("Gudang").Find(&data)
		if tx.Error != nil {
			return user.Core{}, tx.Error
		}
	} else {
		txClient := repo.db.Where("id = ?", userId).Find(&data)
		if txClient.Error != nil {
			return user.Core{}, txClient.Error
		}
	}
	return data.toCore(), nil
}

func (repo *userData) UpdateUser(id int, updateData user.Core) (int, error) {
	dataModel := fromCore(updateData)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *userData) DeleteMitraData(id int) (int, error) {
	var deleteData User
	tx := repo.db.Where("id = ?", id).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}
