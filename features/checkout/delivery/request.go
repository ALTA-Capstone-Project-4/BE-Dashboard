package delivery

import (
	"warehouse/features/checkout"
)

type CheckoutRequest struct {
	NamaBarang       string `json:"nama_barang" form:"nama_barang"`
	MulaiSewa        string `json:"mulai_sewa" form:"mulai_sewa"`
	AkhirSewa        string `json:"akhir_sewa" form:"akhir_sewa"`
	MetodePembayaran string `json:"metode_pembayaran" form:"metode_pembayaran"`
	Status           string `json:"status" form:"status"`
	LahanID          int    `json:"lahan_id" form:"lahan_id"`
}

func toCore(data CheckoutRequest) checkout.Core {
	return checkout.Core{
		NamaBarang:       data.NamaBarang,
		MetodePembayaran: data.MetodePembayaran,
		LahanID:          data.LahanID,
	}
}
