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
	GetMitraId(id int) (Core, error)
	PutMitra(id int, data Core) (int, error)
}

type DataInterface interface {
	AddUser(data Core) (int, error)
	SelectMitra(id int) (Core, error)
	UpdateMitra(id int, data Core) (int, error)
}
