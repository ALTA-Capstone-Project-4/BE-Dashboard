package lahan

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
	Gudang               Gudang
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
}

type DataInterface interface {
	CreateLahan(data Core, user_id int) (int, error)
}
