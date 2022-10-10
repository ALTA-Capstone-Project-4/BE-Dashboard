package usecase

import (
	"errors"
	"testing"
	"warehouse/features/gudang"
	"warehouse/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllGudang(t *testing.T) {
	repo := new(mocks.GudangData)
	dataLahan := []gudang.Lahan{{ID: 1, Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1}}
	dataGudang := gudang.Core{ID: 1, Name: "coco", Latitude: "1234", Longitude: "5678", Location: "jakarta", UserID: 1}
	dataResult := []gudang.Lahan{{ID: 1, Nama: "coco", Luas: "400 m", Panjang: "20 m", Lebar: "30 m", Harga: 200000, Deskripsi: "lahan memanjang", Barang_Tdk_Diizinkan: "makanan", Status: "tidak disewa", FotoLahan: "file.jpg", GudangID: 1, Alamat: "jakarta", Nama_Gudang: "coco"}}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectAllLahan", mock.Anything).Return(dataLahan, nil).Once()
		repo.On("SelectGudangByID", mock.Anything).Return(dataGudang, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetAllGudang(1)
		assert.NoError(t, err)
		assert.Equal(t, dataResult, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data Lahan", func(t *testing.T) {
		repo.On("SelectAllLahan", mock.Anything).Return([]gudang.Lahan{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetAllGudang(1)
		assert.Error(t, err)
		assert.Equal(t, []gudang.Lahan(nil), result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data Gudang", func(t *testing.T) {
		repo.On("SelectAllLahan", mock.Anything).Return(dataLahan, nil).Once()
		repo.On("SelectGudangByID", mock.Anything).Return(gudang.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetAllGudang(1)
		assert.Error(t, err)
		assert.Equal(t, []gudang.Lahan(nil), result)
		repo.AssertExpectations(t)
	})
}

func TestPostGudang(t *testing.T) {
	repo := new(mocks.GudangData)
	input := gudang.Core{Name: "coco", Latitude: "1234", Longitude: "5678", Location: "jakarta"}

	t.Run("create success", func(t *testing.T) {

		repo.On("CreatGudang", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		res, err := usecase.PostGudang(input)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	input_nil := gudang.Core{Name: "", Latitude: "", Longitude: "", Location: "", UserID: 0}
	t.Run("Empty data", func(t *testing.T) {
		usecase := New(repo)
		result, err := usecase.PostGudang(input_nil)

		assert.EqualError(t, err, "tidak boleh ada yang dikosongkan")
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetGudangByID(t *testing.T) {
	repo := new(mocks.GudangData)
	data := gudang.Core{Name: "coco", Latitude: "1234", Longitude: "5678", Location: "jakarta", UserID: 1}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectGudangByID", mock.Anything).Return(data, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetGudangByID(1)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get data", func(t *testing.T) {
		repo.On("SelectGudangByID", mock.Anything).Return(gudang.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetGudangByID(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})

}
