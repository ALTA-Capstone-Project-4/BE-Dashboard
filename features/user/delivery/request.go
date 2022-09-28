package delivery

import "warehouse/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type MitraRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
	FileKTP  string `json:"file_ktp" form:"file_ktp"`
}

func userToCore(data UserRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
	}
}

func mitraToCore(data MitraRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
		Mitra:    user.Mitra{FileKTP: data.FileKTP},
	}
}
