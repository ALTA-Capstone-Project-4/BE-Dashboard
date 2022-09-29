package gudang

type Core struct {
	ID        int
	Name      string
	Photo     string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
}

type UsecaseInterface interface {
	PutGudang(Core) (int, error)
}

type DataInterface interface {
	UpdateGudang(Core) (int, error)
}
