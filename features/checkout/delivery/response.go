package delivery

type LoginResponse struct {
	Token string `json:"token" form:"token"`
	Role  string `json:"role" form:"role"`
	Name  string `json:"name" form:"name"`
}

func FromCore(token, role, name string) LoginResponse {
	return LoginResponse{
		Token: token,
		Role:  role,
		Name:  name,
	}

}
