package usecase

import (
	"warehouse/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostUser(data user.Core) (int, error) {
	if data.Name == "" || data.Email == "" || data.Password == "" || data.Phone == "" || data.Address == "" {
		passByte := []byte(data.Password)
		hashPass, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)

		data.Password = string(hashPass)
	}

	row, err := usecase.userData.AddUser(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *userUsecase) GetUserProfile(id int, userId int) (user.Core, error) {
	data, err := usecase.userData.SelectUserProfile(id, userId)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}

func (usecase *userUsecase) PutUser(id int, updateData user.Core) (int, error) {
	row, err := usecase.userData.UpdateUser(id, updateData)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}

func (usecase *userUsecase) DeleteMitra(id int) (int, error) {
	row, err := usecase.userData.DeleteMitraData(id)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}
