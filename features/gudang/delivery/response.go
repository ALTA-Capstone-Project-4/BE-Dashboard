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

type GudangHomepage struct {
	GudangID  uint   `json:"gudang_id"`
	Ukuran    string `json:"ukuran,omitempty"`
	Harga     int    `json:"harga,omitempty"`
	FotoLahan string `json:"foto_lahan,omitempty"`
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

func fromCore_toHomepage(data gudang.Lahan) GudangHomepage {
	return GudangHomepage{
		GudangID:  uint(data.GudangID),
		Ukuran:    data.Luas,
		Harga:     data.Harga,
		FotoLahan: data.FotoLahan,
	}
}

func fromCoreList(data []gudang.Core) []GudangResponse {
	var dataRes []GudangResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes
}

func fromCore_toHomeList(data []gudang.Lahan) []GudangHomepage {
	var dataRes []GudangHomepage
	for _, v := range data {
		dataRes = append(dataRes, fromCore_toHomepage(v))
	}
	return dataRes
}
