package delivery

import "warehouse/features/lahan"

type LahanRequest struct {
	Nama                 string `json:"nama" form:"nama"`
	Luas                 string `json:"luas" form:"luas"`
	Panjang              string `json:"panjang" form:"panjang"`
	Lebar                string `json:"lebar" form:"lebar"`
	Harga                int    `json:"harga" form:"harga"`
	Deskripsi            string `json:"deskripsi" form:"deskripsi"`
	Fasilitas            string `json:"fasilitas" form:"fasilitas"`
	Barang_Tdk_Diizinkan string `json:"barang_tdk_diizinkan" form:"barang_tdk_diizinkan"`
	FotoLahan            string `json:"foto_lahan" form:"foto_lahan"`
	GudangID             uint   `json:"gudang_id" form:"gudang_id"`
}

func toCore(data LahanRequest) lahan.Core {
	return lahan.Core{
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
	}
}
