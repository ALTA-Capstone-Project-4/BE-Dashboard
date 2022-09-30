package data

import (
	"fmt"
	"warehouse/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func toCore(user User) auth.Core {
	var core = auth.Core{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Status:   user.Status,
	}

	fmt.Println(core.ID)
	fmt.Println(core.Role)
	return core
}
