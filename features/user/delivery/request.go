package delivery

import "warehouse/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Photo    string `json:"photo" form:"photo"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
	FileKTP  string `json:"file_ktp" form:"file_ktp"`
}

func toCore(data UserRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
		Photo:    data.Photo,
		Role:     data.Role,
		Status:   data.Status,
		FileKTP:  data.FileKTP,
	}
}
