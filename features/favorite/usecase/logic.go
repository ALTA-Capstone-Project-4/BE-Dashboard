package usecase

import (
	"warehouse/features/favorite"
)

type favoriteUsecase struct {
	favoriteData favorite.DataInterface
}

func New(data favorite.DataInterface) favorite.UsecaseInterface {
	return &favoriteUsecase{
		favoriteData: data,
	}
}

func (usecase *favoriteUsecase) PostFavorite(data favorite.Core) (int, error) {
	row, err := usecase.favoriteData.AddFavorite(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *favoriteUsecase) GetFavorite(token int) ([]favorite.Core, error) {
	data, err := usecase.favoriteData.SelectFavorite(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}
