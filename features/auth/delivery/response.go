package delivery

type LoginResponse struct {
	Token string `json:"token" form:"token"`
	Role  string `json:"role" form:"role"`
}

func FromCore(token, role string) LoginResponse {
	return LoginResponse{
		Token: token,
		Role:  role,
	}

}
