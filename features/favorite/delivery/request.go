package delivery

import "warehouse/features/favorite"

type FavRequest struct {
	UserID  int
	LahanID int `json:"lahan_id" form:"lahan_id"`
}

func toCore(data FavRequest) favorite.Core {
	return favorite.Core{
		UserID:  uint(data.UserID),
		LahanID: uint(data.LahanID),
	}
}
