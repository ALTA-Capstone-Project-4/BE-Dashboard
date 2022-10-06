package usecase

import (
	"errors"
	"testing"
	"warehouse/features/gudang"
	"warehouse/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestGetAllGudang(t *testing.T) {
// 	repo := new(mocks.GudangData)
// 	dataGudang := []gudang.Core{{Name: "coco", Latitude: "1234", Longitude: "5678", Location: "jakarta", UserID: 1}}

// 	t.Run("Success get data", func(t *testing.T) {
// 		repo.On("SelectAllLahan", mock.Anything).Return(dataGudang, nil).Once()

// 		usecase := New(repo)
// 		result, err := usecase.GetAllGudang(1)
// 		assert.NoError(t, err)
// 		assert.Equal(t, dataGudang, result)
// 		repo.AssertExpectations(t)
// 	})

// t.Run("Failed Get data", func(t *testing.T) {
// 	repo.On("SelectAllLahan", mock.Anything).Return([]gudang.Core{}, errors.New("Error")).Once()

// 	usecase := New(repo)
// 	result, err := usecase.GetAllGudang(1)
// 	assert.Error(t, err)
// 	assert.Equal(t, []gudang.Core([]gudang.Core(nil)), result)
// 	repo.AssertExpectations(t)
// })
// }

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

	// t.Run("Empty data", func(t *testing.T) {
	// 	input := gudang.Core{Name: "", Latitude: "", Longitude: "", Location: "", UserID: 0}

	// 	repo.On("CreatGudang", mock.Anything).Return(-1, errors.New("data tidak boleh kosong")).Once()

	// 	usecase := New(repo)
	// 	result, err := usecase.PostGudang(input)
	// 	assert.EqualError(t, err, "data tidak boleh kosong")
	// 	assert.Equal(t, -1, result)
	// 	repo.AssertExpectations(t)
	// })
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
