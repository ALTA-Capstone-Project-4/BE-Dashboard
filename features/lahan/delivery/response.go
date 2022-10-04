package delivery

import (
	"time"
	"warehouse/features/lahan"
)

type LahanResponse struct {
	ID                   uint      `json:"id"`
	Nama                 string    `json:"nama,omitempty"`
	Luas                 string    `json:"luas,omitempty"`
	Panjang              string    `json:"panjang,omitempty"`
	Lebar                string    `json:"lebar,omitempty"`
	Harga                int       `json:"harga,omitempty"`
	Deskripsi            string    `json:"deskripsi,omitempty"`
	Fasilitas            string    `json:"fasilitas,omitempty"`
	Barang_Tdk_Diizinkan string    `json:"barang_tdk_diizinkan,omitempty"`
	FotoLahan            string    `json:"foto_lahan,omitempty"`
	GudangID             uint      `json:"gudang_id,omitempty"`
	CheckoutFotoBarang   string    `json:"foto_barang,omitempty"`
	CheckoutMulaiSewa    time.Time `json:"mulai_sewa,omitempty"`
	CheckoutAkhirSewa    time.Time `json:"akhir_sewa,omitempty"`
	CheckoutNamaBarang   string    `json:"nama_barang,omitempty"`
	UserName             string    `json:"user_name,omitempty"`
	UserAddress          string    `json:"user_address,omitempty"`
}

func fromCore(data lahan.Core) LahanResponse {
	return LahanResponse{
		ID:                   uint(data.ID),
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
		CheckoutFotoBarang:   data.CheckoutFotoBarang,
		CheckoutMulaiSewa:    data.CheckoutMulaiSewa,
		CheckoutAkhirSewa:    data.CheckoutAkhirSewa,
		CheckoutNamaBarang:   data.CheckoutNamaBarang,
		UserName:             data.UserName,
		UserAddress:          data.UserAddress,
	}
}

func fromCoreList(data []lahan.Core) []LahanResponse {
	var dataRes []LahanResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes
}
