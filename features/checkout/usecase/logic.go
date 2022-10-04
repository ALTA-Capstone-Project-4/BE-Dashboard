package usecase

import (
	"warehouse/features/checkout"
	"warehouse/features/lahan"
)

type checkoutUsecase struct {
	checkoutData checkout.DataInterface
	lahanData    lahan.DataInterface
}

func New(data checkout.DataInterface, dataLahan lahan.DataInterface) checkout.UsecaseInterface {
	return &checkoutUsecase{
		checkoutData: data,
		lahanData:    dataLahan,
	}
}

func (usecase *checkoutUsecase) PostCheckoutByFav(data checkout.Core) (int, error) {

	data.Status = "pending"
	corelahan, errlahan := usecase.lahanData.SelectDetailLahan(data.LahanID, "penitip")
	if errlahan != nil {
		return -1, errlahan
	}
	data.TotalHarga = corelahan.Harga

	row, err := usecase.checkoutData.AddCheckoutByFav(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}
