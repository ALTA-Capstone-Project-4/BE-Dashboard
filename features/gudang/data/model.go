package data

import (
	"warehouse/features/gudang"

	"gorm.io/gorm"
)

type Gudang struct {
	gorm.Model
	Name      string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Photo    string
	FileKTP  string
	Role     string
	Gudang   Gudang
}

func fromCore(dataCore gudang.Core) Gudang {
	dataModel := Gudang{
		Name:      dataCore.Name,
		Latitude:  dataCore.Latitude,
		Longitude: dataCore.Longitude,
		Location:  dataCore.Location,
		UserID:    dataCore.UserID,
	}
	return dataModel
}

func (data *Gudang) toCore() gudang.Core {
	return gudang.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Location:  data.Location,
		UserID:    data.UserID,
	}
}

func toCoreList(data []Gudang) []gudang.Core {
	var dataCore []gudang.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}