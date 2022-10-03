package data

import (
	"warehouse/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Phone    string
	Address  string
	Photo    string
	FileKTP  string
	Role     string
	Status   string
	Gudang   Gudang
	Favorite []Favorite
}

type Favorite struct {
	gorm.Model
	UserID   uint
	LahanID  uint
	GudangID uint
	User     User
}

type Gudang struct {
	gorm.Model
	Name      string
	Location  string
	Latitude  string
	Longitude string
	UserID    uint
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Photo:    dataCore.Photo,
		FileKTP:  dataCore.FileKTP,
		Role:     dataCore.Role,
		Status:   dataCore.Status,
	}
	return dataModel
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:              int(data.ID),
		Name:            data.Name,
		Email:           data.Email,
		Password:        data.Password,
		Phone:           data.Phone,
		Address:         data.Address,
		Photo:           data.Photo,
		FileKTP:         data.FileKTP,
		Role:            data.Role,
		Status:          data.Status,
		GudangName:      data.Gudang.Name,
		GudangLocation:  data.Gudang.Location,
		GudangLatitude:  data.Gudang.Latitude,
		GudangLongitude: data.Gudang.Longitude,
	}
}

func toCoreList(data []User) []user.Core {
	var dataCore []user.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
