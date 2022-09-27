package data

import (
	"warehouse/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	Phone      string
	Address    string
	Role       string
	MitraModel Mitra
}

type Mitra struct {
	ID      int
	FileKTP string
	Status  string
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Phone:      dataCore.Phone,
		Address:    dataCore.Address,
		Role:       dataCore.Role,
		MitraModel: Mitra{FileKTP: dataCore.Mitra.FileKTP},
	}
	return dataModel
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
		Mitra:    user.Mitra{FileKTP: data.MitraModel.FileKTP},
	}
}
