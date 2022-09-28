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
	if data.Email != "" && data.Password != "" {
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

func (usecase *userUsecase) GetMitraId(id int) (user.Core, error) {
	data, err := usecase.userData.SelectMitra(id)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}
