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

	row, err := usecase.lahanData.CreateLahan(data, user_id)

	return row, err

}
