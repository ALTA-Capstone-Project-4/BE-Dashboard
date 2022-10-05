package checkout

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Core struct {
	ID                int
	FotoBarang        string
	NamaBarang        string
	MulaiSewa         time.Time
	AkhirSewa         time.Time
	Periode           int
	MetodePembayaran  string
	Status            string
	TotalHarga        int
	UserID            int
	UserName          string
	LahanID           int
	LahanFoto         string
	LahanNama         string
	LahanHarga        int
	OrderID           string
	TransactionID     string
	BillNumber        string
	TransactionExpire string
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
}

type UsecaseInterface interface {
	CreatePaymentBankTransfer(lahan_id, mitra_id int, reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	GetDataLahan(lahan_id int, role string) (int, int, error)
	PostCheckoutByFav(data Core) (int, error)
	PaymentWebHook(orderID, status string) error
}

type DataInterface interface {
	AddCheckoutByFav(data Core) (int, error)
	CreateDataPayment(reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	SelectPayment(orderID string) (Core, error)
	PaymentDataWebHook(data Core) error
}
