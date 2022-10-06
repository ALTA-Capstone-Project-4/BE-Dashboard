package usecase

import (
	"errors"
	"testing"
	"warehouse/features/lahan"
	"warehouse/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestPostLahan(t *testing.T) {
// 	repo := new(mocks.LahanData)
// 	input := lahan.Core{Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1}

// t.Run("create success", func(t *testing.T) {

// 	repo.On("CreateLahan", mock.Anything, mock.Anything).Return(1, nil).Once()

// 	usecase := New(repo)
// 	res, err := usecase.PostLahan(input, 1)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, res)
// 	repo.AssertExpectations(t)
// })

// t.Run("error add data", func(t *testing.T) {

// 	repo.On("CreateLahan", mock.Anything, mock.Anything).Return(-1, errors.New("there is some error")).Once()

// 	usecase := New(repo)
// 	res, err := usecase.PostLahan(input, 1)
// 	assert.EqualError(t, err, "there is some error")
// 	assert.Equal(t, -1, res)
// 	repo.AssertExpectations(t)
// })
// }

func TestGetDetailLahan(t *testing.T) {
	repo := new(mocks.LahanData)
	data := lahan.Core{Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectDetailLahan", mock.Anything, mock.Anything).Return(data, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetDetailLahan(1, "")
		assert.NoError(t, err)
		assert.Equal(t, data, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get data", func(t *testing.T) {
		repo.On("SelectDetailLahan", mock.Anything, mock.Anything).Return(lahan.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetDetailLahan(1, "")
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestPutLahan(t *testing.T) {
	repo := new(mocks.LahanData)
	newData := lahan.Core{Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1}

	t.Run("Success update data", func(t *testing.T) {
		repo.On("UpdateLahan", mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutLahan(1, 1, newData)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateLahan", mock.Anything, mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()
		newData := lahan.Core{Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1}

		usecase := New(repo)

		result, err := usecase.PutLahan(1, 1, newData)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteLahan(t *testing.T) {
	repo := new(mocks.LahanData)
	data := lahan.Core{Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1}

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteLahan(1, 1, data)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteLahan(1, 1, data)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetLahanClient(t *testing.T) {
	repo := new(mocks.LahanData)
	dataLahanPenitip := []lahan.LahanPenitip{{LahanID: 1, NamaLahan: "Bomby Lahan 1", LuasLahan: "700 m", GudangID: 1, NamaGudang: "Gudang Bomby", AlamatGudang: "bandung", CheckoutID: 1, NamaBarang: "lemari", BillNumber: "39837583645242", StatusBayar: "paid"}}

	t.Run("Success Get data", func(t *testing.T) {
		repo.On("SelectLahan_ByClientID", mock.Anything).Return(dataLahanPenitip, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetLahanClient(1)
		assert.NoError(t, err)
		assert.Equal(t, dataLahanPenitip, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data", func(t *testing.T) {
		repo.On("SelectLahan_ByClientID", mock.Anything).Return([]lahan.LahanPenitip{}, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.GetLahanClient(1)
		assert.Error(t, err)
		assert.Equal(t, []lahan.LahanPenitip([]lahan.LahanPenitip(nil)), result)
		repo.AssertExpectations(t)
	})
}
