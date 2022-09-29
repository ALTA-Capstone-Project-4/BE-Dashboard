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

func (repo *userData) SelectMitraByAdmin(id int) (user.Core, error) {
	var data User

	tx := repo.db.Where("id = ? AND role = ?", id, "mitra").Preload("Gudang").Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *userData) SelectMitra(id int) (user.Core, error) {
	var data User

	tx := repo.db.Where("id = ?", id).Preload("Gudang").Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *userData) UpdateMitra(id int, updateData user.Core) (int, error) {
	dataModel := fromCore(updateData)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

// func (repo *userData) DeleteData(id int, mitra string, client string) (int, error) {
// 	var deleteData User

// 	if deleteData.Role == "admin" {
// 		tx := repo.db.Where("id = ? AND role = ?", id, mitra).Delete(&deleteData)
// 		if tx.Error != nil {
// 			return -1, tx.Error
// 		}
// 	} else if deleteData.Role == "client" {
// 		tx := repo.db.Where("id = ? AND role = ?", id, client).Delete(&deleteData)
// 		if tx.Error != nil {
// 			return -1, tx.Error
// 		}
// 	}

// 	return 1, nil
// }
