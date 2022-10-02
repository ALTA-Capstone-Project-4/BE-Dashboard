package usecase

import (
	"errors"
	"testing"
	"warehouse/features/user"
	"warehouse/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostUser(t *testing.T) {
	repo := new(mocks.UserData)
	input := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("create success", func(t *testing.T) {

		repo.On("AddUser", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		res, err := usecase.PostUser(input)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("error add data", func(t *testing.T) {

		repo.On("AddUser", mock.Anything).Return(-1, errors.New("there is some error")).Once()

		usecase := New(repo)
		res, err := usecase.PostUser(input)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestGetMitraUnverif(t *testing.T) {
	repo := new(mocks.UserData)
	dataUser := []user.Core{{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectMitraUnverif").Return(dataUser, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetMitraUnverif()
		assert.NoError(t, err)
		assert.Equal(t, dataUser, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data", func(t *testing.T) {
		repo.On("SelectMitraUnverif", mock.Anything).Return([]user.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetMitraUnverif()
		assert.Error(t, err)
		assert.Equal(t, []user.Core([]user.Core(nil)), result)
		repo.AssertExpectations(t)
	})
}

func TestPutVerify(t *testing.T) {
	repo := new(mocks.UserData)
	newData := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("Success Update data", func(t *testing.T) {
		repo.On("UpdateVerify", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutVerify(1, newData)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateVerify", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutVerify(1, newData)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetMitraVerified(t *testing.T) {
	repo := new(mocks.UserData)
	dataUser := []user.Core{{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectMitraVerified").Return(dataUser, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetMitraVerified()
		assert.NoError(t, err)
		assert.Equal(t, dataUser, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data", func(t *testing.T) {
		repo.On("SelectMitraVerified", mock.Anything).Return([]user.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetMitraVerified()
		assert.Error(t, err)
		assert.Equal(t, []user.Core([]user.Core(nil)), result)
		repo.AssertExpectations(t)
	})
}

func TestGetMitraByAdmin(t *testing.T) {
	repo := new(mocks.UserData)
	data := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectMitraByAdmin", mock.Anything).Return(data, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetMitraByAdmin(1)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get data", func(t *testing.T) {
		repo.On("SelectMitraByAdmin", mock.Anything).Return(user.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetMitraByAdmin(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetMitra(t *testing.T) {
	repo := new(mocks.UserData)
	data := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectMitra", mock.Anything).Return(data, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetMitra(1)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get data", func(t *testing.T) {
		repo.On("SelectMitra", mock.Anything).Return(user.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetMitra(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestPutMitra(t *testing.T) {
	repo := new(mocks.UserData)
	newData := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("Success update data", func(t *testing.T) {
		repo.On("UpdateMitra", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutMitra(1, newData)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update data", func(t *testing.T) {
		repo.On("UpdateMitra", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutMitra(1, newData)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteMitra(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success delete data.", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteMitra(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete data", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteMitra(1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetClient(t *testing.T) {
	repo := new(mocks.UserData)
	dataClient := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("Success get data", func(t *testing.T) {
		repo.On("SelectClient", mock.Anything).Return(dataClient, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetClient(1)
		assert.NoError(t, err)
		assert.Equal(t, dataClient, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get data", func(t *testing.T) {
		repo.On("SelectClient", mock.Anything).Return(user.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetClient(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestPutClient(t *testing.T) {
	repo := new(mocks.UserData)
	newData := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta", Photo: "photo.jpg", Role: "mitra", Status: "verified", FileKTP: "file_ktp.jpg"}

	t.Run("Success update data", func(t *testing.T) {
		repo.On("UpdateClient", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutClient(1, newData)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update data", func(t *testing.T) {
		repo.On("UpdateClient", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutClient(1, newData)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteClient(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success delete data.", func(t *testing.T) {
		repo.On("DeleteClientData", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteClient(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete data", func(t *testing.T) {
		repo.On("DeleteClientData", mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteClient(1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}
