package delivery

import "warehouse/features/gudang"

type GudangRequest struct {
	Name      string `json:"name" form:"name"`
	Location  string `json:"location" form:"location"`
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
	UserID    int    `json:"user_id" form:"user_id"`
}

func toCore(data GudangRequest) gudang.Core {
	return gudang.Core{
		Name:      data.Name,
		Location:  data.Location,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		UserID:    uint(data.UserID),
	}
}
