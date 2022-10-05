package data

import (
	"warehouse/features/lahan"

	"gorm.io/gorm"
)

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
	Gudang               Gudang
}

type Gudang struct {
	gorm.Model
	Name      string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
	Lahan     []Lahan
}

func fromCore(dataCore lahan.Core) Lahan {
	dataModel := Lahan{
		Nama:                 dataCore.Nama,
		Luas:                 dataCore.Luas,
		Panjang:              dataCore.Panjang,
		Lebar:                dataCore.Lebar,
		Harga:                dataCore.Harga,
		Deskripsi:            dataCore.Deskripsi,
		Fasilitas:            dataCore.Fasilitas,
		Barang_Tdk_Diizinkan: dataCore.Barang_Tdk_Diizinkan,
		FotoLahan:            dataCore.FotoLahan,
		GudangID:             dataCore.GudangID,
		Status:               dataCore.Status,
	}
	return dataModel
}

func (data *Lahan) toCore() lahan.Core {
	return lahan.Core{
		ID:                   int(data.ID),
		Nama:                 data.Nama,
		Luas:                 data.Luas,
		Panjang:              data.Panjang,
		Lebar:                data.Lebar,
		Harga:                data.Harga,
		Deskripsi:            data.Deskripsi,
		Fasilitas:            data.Fasilitas,
		Barang_Tdk_Diizinkan: data.Barang_Tdk_Diizinkan,
		FotoLahan:            data.FotoLahan,
		GudangID:             data.GudangID,
		Status:               data.Status,
		Gudang:               lahan.Gudang{UserID: data.Gudang.UserID},
	}
}

func toCoreList(data []Lahan) []lahan.Core {
	var dataCore []lahan.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
