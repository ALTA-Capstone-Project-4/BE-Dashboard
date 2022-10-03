package data

import (
	"warehouse/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func toCore(user User) auth.Core {
	var core = auth.Core{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Status:   user.Status,
	}

	// fmt.Println(core.ID)
	// fmt.Println(core.Role)
	return core
}
