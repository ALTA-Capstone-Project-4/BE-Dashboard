package data

import (
	"warehouse/features/gudang"

	"gorm.io/gorm"
)

type Gudang struct {
	gorm.Model
	Name      string
	Photo     string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
}

func fromCore(dataCore gudang.Core) Gudang {
	dataModel := Gudang{
		Name:      dataCore.Name,
		Photo:     dataCore.Photo,
		Latitude:  dataCore.Latitude,
		Longitude: dataCore.Longitude,
		Location:  dataCore.Location,
	}
	return dataModel
}

func (data *Gudang) toCore() gudang.Core {
	return gudang.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Photo:     data.Photo,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Location:  data.Location,
	}
}
