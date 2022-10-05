package delivery

import (
	"warehouse/features/lahan"
)

type LahanResponse struct {
	ID                   uint   `json:"id"`
	Nama                 string `json:"nama,omitempty"`
	Luas                 string `json:"luas,omitempty"`
	Panjang              string `json:"panjang,omitempty"`
	Lebar                string `json:"lebar,omitempty"`
	Harga                int    `json:"harga,omitempty"`
	Deskripsi            string `json:"deskripsi,omitempty"`
	Fasilitas            string `json:"fasilitas,omitempty"`
	Barang_Tdk_Diizinkan string `json:"barang_tdk_diizinkan,omitempty"`
	Status               string `json:"status,omitempty"`
	FotoLahan            string `json:"foto_lahan,omitempty"`
	GudangID             uint   `json:"gudang_id,omitempty"`
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
		Status:               data.Status,
		FotoLahan:            data.FotoLahan,
		GudangID:             data.GudangID,
	}
}

func fromCoreList(data []lahan.Core) []LahanResponse {
	var dataRes []LahanResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes
}
