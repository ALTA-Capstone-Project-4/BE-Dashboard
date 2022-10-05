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
	Lahan     []Lahan
	User      User `gorm:"foreignKey:UserID"`
}

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
	Status               string
	FotoLahan            string
	GudangID             uint
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

func (data *Lahan) toLahan() gudang.Lahan {
	return gudang.Lahan{
		ID:                   int(data.ID),
		Nama:                 data.Nama,
		Luas:                 data.Luas,
		Panjang:              data.Panjang,
		Lebar:                data.Lebar,
		Harga:                data.Harga,
		Deskripsi:            data.Deskripsi,
		Fasilitas:            data.Fasilitas,
		Barang_Tdk_Diizinkan: data.Barang_Tdk_Diizinkan,
		Status:               data.Status,
		FotoLahan:            data.FotoLahan,
		GudangID:             int(data.GudangID),
	}
}

func toLahanList(data []Lahan) []gudang.Lahan {
	var dataCore []gudang.Lahan
	for key := range data {
		dataCore = append(dataCore, data[key].toLahan())
	}
	return dataCore
}
