package user

type Core struct {
	ID             int
	Name           string
	Email          string
	Password       string
	Phone          string
	Address        string
	Photo          string
	FileKTP        string
	Role           string
	GudangID       int
	GudangName     string
	GudangLocation string
}

type Gudang struct {
	ID       int
	Name     string
	Location string
}

type UsecaseInterface interface {
	PostUser(data Core) (int, error)
	GetUserProfile(id int, userId int) (Core, error)
	PutUser(id int, data Core) (int, error)
	DeleteMitra(id int) (int, error)
}

type DataInterface interface {
	AddUser(data Core) (int, error)
	SelectUserProfile(id int, userId int) (Core, error)
	UpdateUser(id int, data Core) (int, error)
	DeleteMitraData(id int) (int, error)
}
