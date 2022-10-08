package usecase

import (
	"errors"
	"fmt"
	"testing"
	"time"
	"warehouse/features/checkout"
	"warehouse/features/lahan"
	"warehouse/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostCheckout(t *testing.T) {
	repoCheck := new(mocks.CheckoutData)
	repoLahan := new(mocks.LahanData)
	dataLahan := lahan.Core{
		ID:                   1,
		Nama:                 "my lahan",
		Luas:                 "25m2",
		Panjang:              "5m",
		Lebar:                "5m",
		Harga:                2000000,
		Deskripsi:            "Lahan pribadi",
		Fasilitas:            "Double Lock",
		Barang_Tdk_Diizinkan: "Makanan",
		Status:               "tidak disewa",
		FotoLahan:            "mylahan.png",
		GudangID:             2}

	dataCheck := checkout.Core{
		FotoBarang:        "barang.jpg",
		NamaBarang:        "Motor",
		MulaiSewa:         time.Date(2022, 10, 07, 7, 0, 0, 0, time.Local),
		AkhirSewa:         time.Date(2022, 11, 07, 7, 0, 0, 0, time.Local),
		Periode:           1,
		MetodePembayaran:  "BCA",
		TotalHarga:        2000000,
		Status:            "pending",
		UserID:            2,
		LahanID:           1,
		OrderID:           "order-2022-10-07",
		TransactionID:     "transid-2022-10-07",
		BillNumber:        "83136375605",
		TransactionExpire: "2022-10-07 23:09:55 +0000 UTC"}

	t.Run("Success input data.", func(t *testing.T) {
		repoLahan.On("SelectDetailLahan", dataCheck.LahanID, "penitip").Return(dataLahan, nil).Once()
		repoCheck.On("AddCheckoutByFav", mock.Anything).Return(1, nil).Once()

		usecase := New(repoCheck, repoLahan)
		row, err := usecase.PostCheckoutByFav(dataCheck)
		assert.NoError(t, err)
		assert.Equal(t, 1, row)
		repoCheck.AssertExpectations(t)
		repoLahan.AssertExpectations(t)
	})

	t.Run("Failed input data.", func(t *testing.T) {
		repoLahan.On("SelectDetailLahan", dataCheck.LahanID, "penitip").Return(dataLahan, nil).Once()
		repoCheck.On("AddCheckoutByFav", mock.Anything).Return(1, errors.New("error")).Once()

		usecase := New(repoCheck, repoLahan)
		row, err := usecase.PostCheckoutByFav(dataCheck)
		assert.Error(t, err)
		assert.Equal(t, -1, row)
		repoLahan.AssertExpectations(t)
		repoCheck.AssertExpectations(t)

	})
}

func TestGetDataLahan(t *testing.T) {
	repoCheck := new(mocks.CheckoutData)
	repoLahan := new(mocks.LahanData)
	dataLahan := lahan.Core{
		ID:                   1,
		Nama:                 "my lahan",
		Luas:                 "25m2",
		Panjang:              "5m",
		Lebar:                "5m",
		Harga:                2000000,
		Deskripsi:            "Lahan pribadi",
		Fasilitas:            "Double Lock",
		Barang_Tdk_Diizinkan: "Makanan",
		Status:               "tidak disewa",
		FotoLahan:            "mylahan.png",
		GudangID:             2,
		Gudang:               lahan.Gudang{UserID: 1}}

	t.Run("Success get lahan data.", func(t *testing.T) {
		repoLahan.On("SelectDetailLahan", 1, "penitip").Return(dataLahan, nil).Once()

		usecase := New(repoCheck, repoLahan)
		HargaLahan, user_id, status, err := usecase.GetDataLahan(1, "penitip")
		fmt.Println(HargaLahan, user_id, status)
		assert.NoError(t, err)
		assert.Equal(t, 2000000, HargaLahan)
		assert.Equal(t, 1, user_id)
		assert.Equal(t, "tidak disewa", status)
		repoLahan.AssertExpectations(t)
	})

	t.Run("Fail get lahan data.", func(t *testing.T) {
		repoLahan.On("SelectDetailLahan", 1, "penitip").Return(lahan.Core{}, errors.New("error")).Once()

		usecase := New(repoCheck, repoLahan)
		HargaLahan, user_id, status, err := usecase.GetDataLahan(1, "penitip")
		assert.Error(t, err)
		assert.Equal(t, 0, HargaLahan)
		assert.Equal(t, 0, user_id)
		assert.Equal(t, "", status)
		repoLahan.AssertExpectations(t)
	})
}

func TestPayWebHook(t *testing.T) {
	repoCheck := new(mocks.CheckoutData)
	repoLahan := new(mocks.LahanData)

	dataCheck := checkout.Core{
		FotoBarang:        "barang.jpg",
		NamaBarang:        "Motor",
		MulaiSewa:         time.Date(2022, 10, 07, 7, 0, 0, 0, time.Local),
		AkhirSewa:         time.Date(2022, 11, 07, 7, 0, 0, 0, time.Local),
		Periode:           1,
		MetodePembayaran:  "BCA",
		TotalHarga:        2000000,
		Status:            "pending",
		UserID:            2,
		LahanID:           1,
		OrderID:           "order-2022-10-07",
		TransactionID:     "transid-2022-10-07",
		BillNumber:        "83136375605",
		TransactionExpire: "2022-10-07 23:09:55 +0000 UTC"}

	t.Run("Success payment", func(t *testing.T) {
		repoCheck.On("SelectPayment", dataCheck.OrderID).Return(dataCheck, nil).Once()
		repoCheck.On("PaymentDataWebHook", mock.Anything).Return(nil).Once()

		usecase := New(repoCheck, repoLahan)
		err := usecase.PaymentWebHook("order-2022-10-07", "settlement")
		assert.NoError(t, err)
		repoLahan.AssertExpectations(t)
	})

	t.Run("Payment cancelled", func(t *testing.T) {
		repoCheck.On("SelectPayment", dataCheck.OrderID).Return(dataCheck, nil).Once()
		repoCheck.On("PaymentDataWebHook", mock.Anything).Return(nil).Once()

		usecase := New(repoCheck, repoLahan)
		err := usecase.PaymentWebHook("order-2022-10-07", "cancel")
		assert.NoError(t, err)
		repoLahan.AssertExpectations(t)
	})

	t.Run("Payment denied", func(t *testing.T) {
		repoCheck.On("SelectPayment", dataCheck.OrderID).Return(dataCheck, nil).Once()
		repoCheck.On("PaymentDataWebHook", mock.Anything).Return(nil).Once()

		usecase := New(repoCheck, repoLahan)
		err := usecase.PaymentWebHook("order-2022-10-07", "deny")
		assert.NoError(t, err)
		repoLahan.AssertExpectations(t)
	})

	t.Run("Payment expired", func(t *testing.T) {
		repoCheck.On("SelectPayment", dataCheck.OrderID).Return(dataCheck, nil).Once()
		repoCheck.On("PaymentDataWebHook", mock.Anything).Return(nil).Once()

		usecase := New(repoCheck, repoLahan)
		err := usecase.PaymentWebHook("order-2022-10-07", "expire")
		assert.NoError(t, err)
		repoLahan.AssertExpectations(t)
	})
}
