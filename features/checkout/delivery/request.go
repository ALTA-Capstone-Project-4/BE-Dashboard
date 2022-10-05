package delivery

import (
	"warehouse/features/checkout"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type CheckoutRequest struct {
	NamaBarang       string `json:"nama_barang" form:"nama_barang"`
	MulaiSewa        string `json:"mulai_sewa" form:"mulai_sewa"`
	AkhirSewa        string `json:"akhir_sewa" form:"akhir_sewa"`
	Periode          int    `json:"periode" form:"periode"`
	MetodePembayaran string `json:"metode_pembayaran" form:"metode_pembayaran"`
	LahanID          int    `json:"lahan_id" form:"lahan_id"`
	OrderID          string `json:"orderID" form:"orderID"`
}

func toCore(data CheckoutRequest) checkout.Core {
	return checkout.Core{
		NamaBarang:       data.NamaBarang,
		Periode:          data.Periode,
		MetodePembayaran: data.MetodePembayaran,
		LahanID:          data.LahanID,
	}
}

func ToCoreMidtrans(req checkout.Core) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: int64(req.TotalHarga),
		},
	}
}

type MidtransHookRequest struct {
	TransactionTime   string `form:"transaction_time" json:"transaction_time"`
	TransactionStatus string `form:"transaction_status" json:"transaction_status"`
	OrderID           string `form:"order_id" json:"order_id"`
	MerchantID        string `form:"merchant_id" json:"merchant_id"`
	GrossAmount       string `form:"gross_amount" json:"gross_amount"`
	FraudStatus       string `form:"fraud_status" json:"fraud_status"`
	Currency          string `form:"currency" json:"currency"`
}
