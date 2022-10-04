package data

import (
	"time"
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
	FotoLahan            string
	GudangID             uint
	Checkout             Checkout
	Gudang               Gudang
}

// type Favorite struct {
// 	gorm.Model
// 	UserID  uint
// 	LahanID uint
// 	User    User
// }

type Checkout struct {
	gorm.Model
	FotoBarang       string
	NamaBarang       string
	MulaiSewa        time.Time
	AkhirSewa        time.Time
	MetodePembayaran string
	Status           string
	TotalHarga       int
	UserID           int
	LahanID          int
	User             User
}

// type Gudang struct {
// 	gorm.Model
// 	Name      string
// 	Latitude  string
// 	Longitude string
// 	Location  string
// 	UserID    uint
// 	Lahan     []Lahan
// }

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Photo    string
	FileKTP  string
	Role     string
	Status   string
	Lahan    []Lahan
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
		CheckoutFotoBarang:   data.Checkout.FotoBarang,
		CheckoutMulaiSewa:    data.Checkout.MulaiSewa,
		CheckoutAkhirSewa:    data.Checkout.AkhirSewa,
		CheckoutNamaBarang:   data.Checkout.NamaBarang,
		// UserName:             data.User.Name,
		// UserAddress:          data.User.Address,
	}
}

func toCoreList(data []Lahan) []lahan.Core {
	var dataCore []lahan.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
