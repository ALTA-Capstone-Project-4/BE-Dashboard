package lahan

import "time"

type Core struct {
	ID                   int
	Nama                 string
	Luas                 string
	Panjang              string
	Lebar                string
	Harga                int
	Deskripsi            string
	Fasilitas            string
	Barang_Tdk_Diizinkan string
	FotoLahan            string
	GudangID             uint
	CheckoutFotoBarang   string
	CheckoutMulaiSewa    time.Time
	CheckoutAkhirSewa    time.Time
	CheckoutNamaBarang   string
	UserName             string
	UserAddress          string
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

type Gudang struct {
	ID        int
	Name      string
	Latitude  string
	Longitude string
	Location  string
	UserID    uint
}

type UsecaseInterface interface {
	PostLahan(data Core, user_id int) (int, error)
	PutLahan(id int, data Core) (int, error)
	GetDetailLahan(id int, role string) (Core, error)
	DeleteLahan(token int, id int) (int, error)
}

type DataInterface interface {
	CreateLahan(data Core, user_id int) (int, error)
	UpdateLahan(id int, data Core) (int, error)
	SelectDetailLahan(id int, role string) (Core, error)
	DeleteData(token int, id int) (int, error)
}
