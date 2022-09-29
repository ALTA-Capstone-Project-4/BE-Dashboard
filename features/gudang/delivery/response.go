package delivery

import (
	"warehouse/features/gudang"
)

type GudangResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name,omitempty"`
	Location  string `json:"location,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
}

func fromCore(data gudang.Core) GudangResponse {
	return GudangResponse{
		ID:        uint(data.ID),
		Name:      data.Name,
		Location:  data.Location,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		UserID:    int(data.UserID),
	}
}

func fromCoreList(data []gudang.Core) []GudangResponse {
	var dataRes []GudangResponse
	for _, v := range data {
		dataRes = append(dataRes, GudangResponse{
			ID:        uint(v.ID),
			Name:      v.Name,
			Location:  v.Location,
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
			UserID:    int(v.UserID),
		})
	}

	return dataRes
}
