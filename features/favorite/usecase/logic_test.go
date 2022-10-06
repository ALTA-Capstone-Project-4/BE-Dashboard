package usecase

import (
	"errors"
	"testing"
	"warehouse/features/favorite"
	"warehouse/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostFavorite(t *testing.T) {
	repo := new(mocks.FavoriteData)
	dataInput := favorite.Core{UserID: 1, LahanID: 3}

	t.Run("Success Insert data.", func(t *testing.T) {
		repo.On("AddFavorite", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		result, err := usecase.PostFavorite(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Insert data.", func(t *testing.T) {
		repo.On("AddFavorite", mock.Anything).Return(1, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.PostFavorite(dataInput)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetFavorite(t *testing.T) {
	repo := new(mocks.FavoriteData)
	dataProduct := []favorite.Core{{UserID: 1, LahanID: 3}}

	t.Run("Success Get all data.", func(t *testing.T) {
		repo.On("SelectFavorite", mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetFavorite(1)
		assert.NoError(t, err)
		assert.Equal(t, dataProduct, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get all data.", func(t *testing.T) {
		repo.On("SelectFavorite", mock.Anything).Return([]favorite.Core{}, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.GetFavorite(1)
		assert.Error(t, err)
		assert.Equal(t, []favorite.Core([]favorite.Core(nil)), result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteFavorite(t *testing.T) {
	repo := new(mocks.FavoriteData)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteFavorite(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteFavorite(1, 1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}
