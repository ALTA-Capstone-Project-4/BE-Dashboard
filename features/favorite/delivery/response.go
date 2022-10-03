package delivery

import "warehouse/features/favorite"

type FavResponse struct {
	ID             uint
	UserName       string
	LahanName      string
	LahanHarga     int
	LahanFotoLahan string
}

func fromCore(data favorite.Core) FavResponse {
	return FavResponse{
		ID:             uint(data.ID),
		UserName:       data.UserName,
		LahanName:      data.LahanName,
		LahanHarga:     data.LahanHarga,
		LahanFotoLahan: data.LahanFotoLahan,
	}
}

func fromCoreList(data []favorite.Core) []FavResponse {

	var dataRes []FavResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes
}
