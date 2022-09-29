package delivery

import "warehouse/features/gudang"

type GudangRequest struct {
	Name     string `json:"name" form:"name"`
	Photo    string `json:"photo" form:"photo"`
	Location string `json:"location" form:"location"`
}

func toCore(data GudangRequest) gudang.Core {
	return gudang.Core{
		Name:     data.Name,
		Photo:    data.Photo,
		Location: data.Location,
	}
}
