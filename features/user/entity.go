package user

type Core struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Role     string
	Mitra    Mitra
}

type Mitra struct {
	ID      int
	FileKTP string
	Status  string
}

type UsecaseInterface interface {
}

type DataInterface interface {
}
