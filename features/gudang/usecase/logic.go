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

func (usecase *gudangUsecase) GetAllGudang(page int) ([]gudang.Lahan, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 8

	data, err := usecase.gudangData.SelectAllLahan(offset)
	if err != nil {
		return nil, err
	}

	for key := range data {
		dataGudang, err := usecase.gudangData.SelectGudangByID(data[key].GudangID)

		if err != nil {
			return nil, err
		}

		data[key].Alamat = dataGudang.Location
		data[key].Nama_Gudang = dataGudang.Name
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
