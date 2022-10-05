package delivery

import (
	"warehouse/features/gudang"
)

type GudangResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name,omitempty"`
	Location  string         `json:"address,omitempty"`
	Latitude  string         `json:"latitude,omitempty"`
	Longitude string         `json:"longitude,omitempty"`
	UserID    int            `json:"user_id,omitempty"`
	Lahan     []gudang.Lahan `json:"lahan,omitempty"`
}

type DataHome struct {
	GudangID             int    `json:"gudang_id,omitempty"`
	ID                   uint   `json:"id_lahan"`
	Nama_Gudang          string `json:"nama_gudang,omitempty"`
	Nama                 string `json:"nama,omitempty"`
	Luas                 string `json:"luas,omitempty"`
	Panjang              string `json:"panjang,omitempty"`
	Lebar                string `json:"lebar,omitempty"`
	Harga                int    `json:"harga,omitempty"`
	Deskripsi            string `json:"deskripsi,omitempty"`
	Fasilitas            string `json:"fasilitas,omitempty"`
	Barang_Tdk_Diizinkan string `json:"barang_tdk_diizinkan,omitempty"`
	FotoLahan            string `json:"foto_lahan,omitempty"`
	Alamat               string `json:"address,omitempty"`
}

func fromCore(data gudang.Core) GudangResponse {
	return GudangResponse{
		ID:        uint(data.ID),
		Name:      data.Name,
		Location:  data.Location,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		UserID:    int(data.UserID),
		Lahan:     data.Lahan,
	}
}

func fromCoreList(data []gudang.Core) []GudangResponse {
	var dataRes []GudangResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes
}

func fromCore_toHome(data gudang.Lahan) DataHome {
	return DataHome{
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
		Alamat:               data.Alamat,
		Nama_Gudang:          data.Nama_Gudang,
	}
}

func fromCore_toHomeList(data []gudang.Lahan) []DataHome {
	var dataRes []DataHome
	for _, v := range data {
		dataRes = append(dataRes, fromCore_toHome(v))
	}
	return dataRes
}
