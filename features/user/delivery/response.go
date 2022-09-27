package delivery

import "warehouse/features/user"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
	Role    string `json:"role,omitempty"`
}

type MitraResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
	Role    string `json:"role,omitempty"`
	FileKTP string `json:"file_ktp" form:"file_ktp"`
}

func fromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:      uint(data.ID),
		Name:    data.Name,
		Phone:   data.Phone,
		Email:   data.Email,
		Address: data.Address,
		Role:    data.Role,
	}
}

func fromCoreMitra(data user.Core) MitraResponse {
	return MitraResponse{
		ID:      uint(data.ID),
		Name:    data.Name,
		Phone:   data.Phone,
		Email:   data.Email,
		Address: data.Address,
		Role:    data.Role,
		FileKTP: data.Mitra.FileKTP,
	}
}
