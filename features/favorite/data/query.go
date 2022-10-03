package data

import (
	"fmt"
	"warehouse/features/favorite"
	modelLahan "warehouse/features/lahan/data"

	"gorm.io/gorm"
)

type favoriteData struct {
	db *gorm.DB
}

func New(db *gorm.DB) favorite.DataInterface {
	return &favoriteData{
		db: db,
	}
}

func (repo *favoriteData) AddFavorite(data favorite.Core) (int, error) {
	var lahan modelLahan.Lahan
	repo.db.Where("id = ?", data.LahanID).Find(&lahan)
	data.LahanID = lahan.ID
	fmt.Println(data)
	// var dbCek Lahan
	// repo.db.First(&dbCek, "lahan_id = ?", data.LahanID)

	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *favoriteData) SelectFavorite(token int) ([]favorite.Core, error) {
	var data []Favorite
	tx := repo.db.Model(&Favorite{}).Where("user_id = ?", token).Preload("Lahan").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
}
