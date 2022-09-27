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
	gorm.Model
	FileKTP string
	Status  string
	UserID  int
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
	}
}
