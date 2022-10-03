package favorite

type Core struct {
	ID             int
	UserID         uint
	UserName       string
	LahanID        uint
	LahanName      string
	LahanHarga     int
	LahanFotoLahan string
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
}

type UsecaseInterface interface {
	PostFavorite(data Core) (int, error)
}

type DataInterface interface {
	AddFavorite(data Core) (int, error)
}
