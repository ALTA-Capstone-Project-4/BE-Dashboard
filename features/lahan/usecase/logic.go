package usecase

import (
	"warehouse/features/lahan"
)

type lahanUsecase struct {
	lahanData lahan.DataInterface
}

func New(data lahan.DataInterface) lahan.UsecaseInterface {
	return &lahanUsecase{
		lahanData: data,
	}
}
