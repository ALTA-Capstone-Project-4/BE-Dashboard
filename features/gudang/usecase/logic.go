package usecase

import (
	"errors"
	"warehouse/features/gudang"
)

type gudangUsecase struct {
	gudangData gudang.DataInterface
}

func New(data gudang.DataInterface) gudang.UsecaseInterface {
	return &gudangUsecase{
		gudangData: data,
	}
}

func (usecase *gudangUsecase) PutGudang(id int, data gudang.Core) (int, error) {
	row, err := usecase.gudangData.UpdateGudang(id, data)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}

func (usecase *gudangUsecase) GetAllGudang(page int) ([]gudang.Lahan, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 8

	data, err := usecase.gudangData.SelectAllGudang(offset)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (usecase *gudangUsecase) PostGudang(data gudang.Core) (int, error) {
	if data.Name == "" || data.Location == "" || data.Latitude == "" || data.Longitude == "" {
		return -1, errors.New("tidak boleh ada yang dikosongkan")
	}

	row, err := usecase.gudangData.CreatGudang(data)

	return row, err
}

func (usecase *gudangUsecase) GetGudangByID(gudang_id int) (gudang.Core, error) {
	data, err := usecase.gudangData.SelectGudangByID(gudang_id)

	if err != nil {
		return gudang.Core{}, err
	}

	return data, nil
}
