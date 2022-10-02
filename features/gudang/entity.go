package gudang

type Core struct {
	ID        int
	Name      string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Photo    string
	FileKTP  string
	Role     string
}

type UsecaseInterface interface {
	PutGudang(id int, data Core) (int, error)
	GetAllGudang() ([]Core, error)
	PostGudang(data Core) (int, error)
}

type DataInterface interface {
	UpdateGudang(id int, data Core) (int, error)
	SelectAllGudang() ([]Core, error)
	CreatGudang(data Core) (int, error)
}
