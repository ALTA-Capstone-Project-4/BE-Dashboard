package data

import (
	"warehouse/features/favorite"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint
	LahanID uint
	User    User  `gorm:"foreignKey:UserID"`
	Lahan   Lahan `gorm:"foreignKey:LahanID"`
}

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	Photo     string
	FileKTP   string
	Role      string
	Status    string
	Favorites []Favorite
}

type Lahan struct {
	gorm.Model
	Nama                 string
	Luas                 string
	Panjang              string
	Lebar                string
	Harga                int
	Deskripsi            string
	Fasilitas            string
	Barang_Tdk_Diizinkan string
	FotoLahan            string
	GudangID             uint
	Favorites            []Favorite
}

func fromCore(dataCore favorite.Core) Favorite {
	dataModel := Favorite{
		UserID:  dataCore.UserID,
		LahanID: dataCore.LahanID,
	}
	return dataModel
}

func (data *Favorite) toCore() favorite.Core {
	return favorite.Core{
		ID:             int(data.ID),
		UserID:         data.UserID,
		LahanID:        data.LahanID,
		LahanName:      data.Lahan.Nama,
		LahanHarga:     data.Lahan.Harga,
		LahanFotoLahan: data.Lahan.FotoLahan,
	}
}

func toCoreList(data []Favorite) []favorite.Core {
	var dataCore []favorite.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
