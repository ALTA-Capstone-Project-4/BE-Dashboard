package auth

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
	Status   string
}

type UsecaseInterface interface {
	LoginAuthorized(email, password string) (string, string, string)
}

type DataInterface interface {
	LoginUser(email string) (Core, error)
}
