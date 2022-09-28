package delivery

import "warehouse/features/auth"

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func toCore(data AuthRequest) auth.Core {
	return auth.Core{
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
