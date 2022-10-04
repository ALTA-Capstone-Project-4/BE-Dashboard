package checkout

import "time"

type Core struct {
	ID               int
	FotoBarang       string
	NamaBarang       string
	MulaiSewa        time.Time
	AkhirSewa        time.Time
	MetodePembayaran string
	Status           string
	TotalHarga       int
	UserID           int
	UserName         string
	LahanID          int
	LahanFoto        string
	LahanNama        string
	LahanHarga       int
	FavoriteID       int
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
	UserID               uint
	CheckoutID           uint
}

type UsecaseInterface interface {
	PostCheckoutByFav(data Core) (int, error)
}

type DataInterface interface {
	AddCheckoutByFav(data Core) (int, error)
}
