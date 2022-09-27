package user

type Core struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Role     string
	MitraKTP string
}

type Mitra struct {
	ID      int
	FileKTP string
	Status  string
	UserID  int
}

type UsecaseInterface interface {
}

type DataInterface interface {
}
