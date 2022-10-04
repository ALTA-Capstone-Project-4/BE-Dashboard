package data

import (
	"time"
	"warehouse/features/checkout"

	"gorm.io/gorm"
)

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
	OrderID          string
	TransactionID    string
	User             User  `gorm:"foreignKey:UserID"`
	Lahan            Lahan `gorm:"foreignKey:LahanID"`
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
	FotoLahan            string
	GudangID             uint
	UserID               uint
}

func toCore(data Checkout) checkout.Core {
	var core = checkout.Core{
		ID:               int(data.ID),
		FotoBarang:       data.FotoBarang,
		NamaBarang:       data.NamaBarang,
		MulaiSewa:        data.MulaiSewa,
		AkhirSewa:        data.AkhirSewa,
		MetodePembayaran: data.MetodePembayaran,
		Status:           data.Status,
		TotalHarga:       data.TotalHarga,
		UserID:           data.UserID,
		UserName:         data.User.Name,
		LahanID:          data.LahanID,
		LahanFoto:        data.Lahan.FotoLahan,
		LahanNama:        data.Lahan.Nama,
		LahanHarga:       data.Lahan.Harga,
		OrderID:          data.OrderID,
		TransactionID:    data.TransactionID,
	}
	return core
}

func fromCore(dataCore checkout.Core) Checkout {
	dataModel := Checkout{
		FotoBarang:       dataCore.FotoBarang,
		NamaBarang:       dataCore.NamaBarang,
		MulaiSewa:        dataCore.MulaiSewa,
		AkhirSewa:        dataCore.AkhirSewa,
		MetodePembayaran: dataCore.MetodePembayaran,
		Status:           dataCore.Status,
		TotalHarga:       dataCore.TotalHarga,
		LahanID:          dataCore.LahanID,
		UserID:           dataCore.UserID,
		OrderID:          dataCore.OrderID,
		TransactionID:    dataCore.TransactionID,
	}
	return dataModel
}

func (data *Checkout) toCoreMidtrans() checkout.Core {
	return checkout.Core{
		ID:               int(data.ID),
		FotoBarang:       data.FotoBarang,
		NamaBarang:       data.NamaBarang,
		MulaiSewa:        data.MulaiSewa,
		AkhirSewa:        data.AkhirSewa,
		MetodePembayaran: data.MetodePembayaran,
		Status:           data.Status,
		TotalHarga:       data.TotalHarga,
		UserID:           data.UserID,
		UserName:         data.User.Name,
		LahanID:          data.LahanID,
		LahanFoto:        data.Lahan.FotoLahan,
		LahanNama:        data.Lahan.Nama,
		LahanHarga:       data.Lahan.Harga,
		OrderID:          data.OrderID,
		TransactionID:    data.TransactionID,
	}
}
