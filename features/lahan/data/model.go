package data

import (
	"time"
	modelCheckout "warehouse/features/checkout/data"
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
	Gudang               Gudang `gorm:"foreignKey:GudangID"`
	Checkout             modelCheckout.Checkout
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

type Checkout struct {
	gorm.Model
	FotoBarang        string
	NamaBarang        string
	MulaiSewa         time.Time
	AkhirSewa         time.Time
	Periode           int
	MetodePembayaran  string
	Status            string
	TotalHarga        int
	UserID            int
	LahanID           int
	OrderID           string
	TransactionID     string
	BillNumber        string
	TransactionExpire string
	Lahan             Lahan
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

func fromCheckout_toLahan(data []Checkout) []Lahan {
	var dataCore []Lahan
	for key := range data {
		dataCore = append(dataCore, data[key].Lahan)
	}
	return dataCore
}

func toLahanPenitip_FromCheckout(dataCheckout Checkout) lahan.LahanPenitip {
	dataModel := lahan.LahanPenitip{
		CheckoutID:  int(dataCheckout.ID),
		NamaBarang:  dataCheckout.NamaBarang,
		BillNumber:  dataCheckout.BillNumber,
		StatusBayar: dataCheckout.Status,
		MulaiSewa:   dataCheckout.MulaiSewa,
		AkhirSewa:   dataCheckout.AkhirSewa,
	}

	return dataModel
}

func toLahanPenitipList(dataCheckout []Checkout, dataLahan []Lahan, dataGudang []Gudang) []lahan.LahanPenitip {
	var dataCore []lahan.LahanPenitip

	for _, v := range dataCheckout {
		dataCore = append(dataCore, toLahanPenitip_FromCheckout(v))
	}

	for key := 0; key < len(dataCore); key++ {

		dataCore[key].LahanID = int(dataLahan[key].ID)
		dataCore[key].NamaLahan = dataLahan[key].Nama
		dataCore[key].LuasLahan = dataLahan[key].Panjang + " x " + dataLahan[key].Lebar
		dataCore[key].GudangID = int(dataGudang[key].ID)
		dataCore[key].NamaGudang = dataGudang[key].Name
		dataCore[key].AlamatGudang = dataGudang[key].Location

	}

	return dataCore
}
