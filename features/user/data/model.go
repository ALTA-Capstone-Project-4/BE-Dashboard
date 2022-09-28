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
	GudangID int
	Gudang   Gudang
}

type Gudang struct {
	gorm.Model
	Name     string
	Location string
	User     []User
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
		GudangID: dataCore.GudangID,
	}
	return dataModel
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:             int(data.ID),
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		Phone:          data.Phone,
		FileKTP:        data.FileKTP,
		Address:        data.Address,
		GudangID:       data.GudangID,
		GudangName:     data.Gudang.Name,
		GudangLocation: data.Gudang.Location,
	}
}
