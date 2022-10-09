package delivery

import (
	"time"
	"warehouse/features/lahan"
)

type LahanResponse struct {
	ID                   uint   `json:"id"`
	Nama                 string `json:"nama,omitempty"`
	Luas                 string `json:"luas,omitempty"`
	Panjang              string `json:"panjang,omitempty"`
	Lebar                string `json:"lebar,omitempty"`
	Harga                int    `json:"harga,omitempty"`
	Deskripsi            string `json:"deskripsi,omitempty"`
	Fasilitas            string `json:"fasilitas,omitempty"`
	Barang_Tdk_Diizinkan string `json:"barang_tdk_diizinkan,omitempty"`
	Status               string `json:"status,omitempty"`
	FotoLahan            string `json:"foto_lahan,omitempty"`
	GudangID             uint   `json:"gudang_id,omitempty"`
}

type LahanPenitip struct {
	CheckoutID   int       `json:"checkout_id,omitempty"`
	NamaBarang   string    `json:"nama_barang,omitempty"`
	FotoBarang   string    `json:"foto_barang,omitempty"`
	BillNumber   string    `json:"bill_number,omitempty"`
	StatusBayar  string    `json:"status_pembayaran,omitempty"`
	MulaiSewa    time.Time `json:"mulai_sewa,omitempty"`
	AkhirSewa    time.Time `json:"akhir_sewa,omitempty"`
	LahanID      int       `json:"lahan_id,omitempty"`
	NamaLahan    string    `json:"nama_lahan,omitempty"`
	LuasLahan    string    `json:"luas_lahan,omitempty"`
	GudangID     int       `json:"gudang_id,omitempty"`
	NamaGudang   string    `json:"nama_gudang,omitempty"`
	AlamatGudang string    `json:"alamat_gudang,omitempty"`
}

func fromCore(data lahan.Core) LahanResponse {
	return LahanResponse{
		ID:                   uint(data.ID),
		Nama:                 data.Nama,
		Luas:                 data.Luas,
		Panjang:              data.Panjang,
		Lebar:                data.Lebar,
		Harga:                data.Harga,
		Deskripsi:            data.Deskripsi,
		Fasilitas:            data.Fasilitas,
		Barang_Tdk_Diizinkan: data.Barang_Tdk_Diizinkan,
		Status:               data.Status,
		FotoLahan:            data.FotoLahan,
		GudangID:             data.GudangID,
	}
}

func fromCoreList(data []lahan.Core) []LahanResponse {
	var dataRes []LahanResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes
}

func fromLahanPenitip(data lahan.LahanPenitip) LahanPenitip {
	return LahanPenitip{
		CheckoutID:   data.CheckoutID,
		NamaBarang:   data.NamaBarang,
		FotoBarang:   data.FotoBarang,
		BillNumber:   data.BillNumber,
		StatusBayar:  data.StatusBayar,
		MulaiSewa:    data.MulaiSewa,
		AkhirSewa:    data.AkhirSewa,
		LahanID:      data.LahanID,
		NamaLahan:    data.NamaLahan,
		LuasLahan:    data.LuasLahan,
		GudangID:     data.GudangID,
		NamaGudang:   data.NamaGudang,
		AlamatGudang: data.AlamatGudang,
	}
}

func fromLahanPenitipList(data []lahan.LahanPenitip) []LahanPenitip {
	var dataRes []LahanPenitip
	for _, v := range data {
		dataRes = append(dataRes, fromLahanPenitip(v))
	}

	return dataRes
}
