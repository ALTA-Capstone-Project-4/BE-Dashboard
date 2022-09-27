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
