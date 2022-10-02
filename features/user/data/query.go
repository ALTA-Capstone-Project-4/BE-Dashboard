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

func (repo *userData) SelectMitraUnverif() ([]user.Core, error) {
	var data []User
	tx := repo.db.Model(&User{}).Where("status = ?", "unverified").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
}

func (repo *userData) UpdateVerify(id int, status user.Core) (int, error) {
	dataModel := fromCore(status)

	tx := repo.db.Model(&User{}).Where("id = ? AND status = ?", id, "unverified").Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *userData) SelectMitraVerified() ([]user.Core, error) {
	var data []User
	tx := repo.db.Model(&User{}).Where("status = ?", "verified").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
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

	tx := repo.db.Where("id = ? AND role = ?", id, "mitra").Preload("Gudang").Find(&data)
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

func (repo *userData) DeleteData(id int) (int, error) {
	var deleteData User

	tx := repo.db.Where("id = ? AND role = ?", id, "mitra").Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}

	return 1, nil
}

func (repo *userData) SelectClient(id int) (user.Core, error) {
	var data User

	tx := repo.db.Where("id = ? AND role = ?", id, "penitip").Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *userData) UpdateClient(id int, updateData user.Core) (int, error) {
	dataModel := fromCore(updateData)

	if dataModel.Photo == "" || dataModel.Photo != "" {
		tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataModel)
		if tx.Error != nil {
			return -1, tx.Error
		}
		return 1, nil
	}
	return 1, nil
}

func (repo *userData) DeleteClientData(id int) (int, error) {
	var deleteData User

	tx := repo.db.Where("id = ?", id).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}

	return 1, nil
}
