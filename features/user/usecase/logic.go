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

	passByte := []byte(data.Password)
	hashPass, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)

	data.Password = string(hashPass)

	row, err := usecase.userData.AddUser(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *userUsecase) GetMitraUnverif() ([]user.Core, error) {
	data, err := usecase.userData.SelectMitraUnverif()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *userUsecase) PutVerify(id int, status user.Core) (int, error) {
	row, err := usecase.userData.UpdateVerify(id, status)
	if err != nil || row < 1 {
		return -1, err

	}

	return 1, nil
}

func (usecase *userUsecase) GetMitraVerified() ([]user.Core, error) {
	data, err := usecase.userData.SelectMitraVerified()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *userUsecase) GetMitraByAdmin(id int) (user.Core, error) {
	data, err := usecase.userData.SelectMitraByAdmin(id)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}

func (usecase *userUsecase) GetMitra(id int) (user.Core, error) {
	data, err := usecase.userData.SelectMitra(id)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}

func (usecase *userUsecase) PutMitra(id int, updateData user.Core) (int, error) {
	passByte := []byte(updateData.Password)
	hashPass, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)

	updateData.Password = string(hashPass)

	row, err := usecase.userData.UpdateMitra(id, updateData)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}

func (usecase *userUsecase) DeleteMitra(id int) (int, error) {
	row, err := usecase.userData.DeleteData(id)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}

func (usecase *userUsecase) GetClient(id int) (user.Core, error) {
	data, err := usecase.userData.SelectClient(id)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}

func (usecase *userUsecase) PutClient(id int, updateData user.Core) (int, error) {
	passByte := []byte(updateData.Password)
	hashPass, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)

	updateData.Password = string(hashPass)

	row, err := usecase.userData.UpdateClient(id, updateData)
	if err != nil || row < 1 {
		return -1, err
	}
	return 1, nil

}

func (usecase *userUsecase) DeleteClient(id int) (int, error) {
	row, err := usecase.userData.DeleteClientData(id)
	if err != nil || row < 1 {
		return -1, err
	}

	return 1, nil
}
