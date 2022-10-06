package usecase

import (
	"errors"
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

func (usecase *lahanUsecase) PostLahan(data lahan.Core, user_id int) (int, error) {
	if data.Nama == "" || data.Luas == "" || data.Panjang == "" || data.Lebar == "" || data.Harga == 0 || data.Deskripsi == "" || data.Fasilitas == "" || data.Barang_Tdk_Diizinkan == "" || data.FotoLahan == "" || user_id == 0 {
		return -1, errors.New("tidak boleh ada yang dikosongkan")
	}

	data.Status = "tidak disewa"

	row, err := usecase.lahanData.CreateLahan(data, user_id)

	return row, err
}

func (usecase *lahanUsecase) GetDetailLahan(id int, role string) (lahan.Core, error) {
	data, err := usecase.lahanData.SelectDetailLahan(id, role)
	if err != nil {
		return lahan.Core{}, err
	}

	return data, nil
}

func (usecase *lahanUsecase) PutLahan(id int, token int, data lahan.Core) (int, error) {
	row, err := usecase.lahanData.UpdateLahan(id, token, data)
	if err != nil || row < 1 {
		return -1, err
	}
	return 1, nil
}

func (usecase *lahanUsecase) DeleteLahan(id int, token int, data lahan.Core) (int, error) {
	row, err := usecase.lahanData.DeleteData(id, token, data)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}

func (usecase *lahanUsecase) GetLahanClient(token int) ([]lahan.LahanPenitip, error) {
	data, err := usecase.lahanData.SelectLahan_ByClientID(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}
