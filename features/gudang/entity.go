package gudang

type Core struct {
	ID        int
	Name      string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
	Lahan     []Lahan
	User      User
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
	Status   string
}

type Lahan struct {
	ID                   int
	Nama_Gudang          string
	Nama                 string
	Luas                 string
	Panjang              string
	Lebar                string
	Harga                int
	Deskripsi            string
	Fasilitas            string
	Barang_Tdk_Diizinkan string
	Status               string
	FotoLahan            string
	GudangID             int
	Alamat               string
}

type UsecaseInterface interface {
	GetAllGudang(page int) ([]Lahan, error)
	PostGudang(data Core) (int, error)
	GetGudangByID(gudang_id int) (Core, error)
	UpdateGudangMitra(user_id int, data Core) (int, error)
}

type DataInterface interface {
	SelectAllLahan(offset int) ([]Lahan, error)
	CreatGudang(data Core) (int, error)
	SelectGudangByID(gudang_id int) (Core, error)
	UpdateGudang(user_id int, data Core) (int, error)
}
