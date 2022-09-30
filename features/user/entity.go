package user

type Core struct {
	ID              int
	Name            string
	Email           string
	Password        string
	Phone           string
	Address         string
	FileKTP         string
	Role            string
	Status          string
	Photo           string
	GudangID        int
	GudangName      string
	GudangLocation  string
	GudangPhoto     string
	GudangLatitude  string
	GudangLongitude string
}

type Gudang struct {
	ID        int
	Name      string
	Photo     string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
}

type UsecaseInterface interface {
	PostUser(data Core) (int, error)
	GetMitraByAdmin(id int) (Core, error)
	GetMitra(id int) (Core, error)
	PutMitra(id int, data Core) (int, error)
	DeleteMitra(id int) (int, error)
	GetClient(id int) (Core, error)
	PutClient(id int, data Core) (int, error)
	DeleteClient(id int) (int, error)
}

type DataInterface interface {
	AddUser(data Core) (int, error)
	SelectMitraByAdmin(id int) (Core, error)
	SelectMitra(id int) (Core, error)
	UpdateMitra(id int, data Core) (int, error)
	DeleteData(id int) (int, error)
	SelectClient(id int) (Core, error)
	UpdateClient(id int, data Core) (int, error)
	DeleteClientData(id int) (int, error)
}
