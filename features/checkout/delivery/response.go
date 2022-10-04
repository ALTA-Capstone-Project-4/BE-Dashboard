package delivery

import (
	"time"
	"warehouse/config"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Payment struct {
	OrderID           string    `json:"orderID" form:"orderID"`
	TransactionID     string    `json:"transactionID" form:"transactionID"`
	PaymentMethod     string    `json:"paymentMethod" form:"paymentMethod"`
	BillNumber        string    `json:"billNumber" form:"billNumber"`
	Bank              string    `json:"bank" form:"bank"`
	GrossAmount       string    `json:"grossAmount" form:"grossAmount"`
	TransactionTime   time.Time `json:"transactionTime" form:"transactionTime"`
	TransactionExpire time.Time `json:"transactionExpired" form:"transactionExpired"`
	TransactionStatus string    `json:"transactionStatus" form:"transactionStatus"`
}

func FromMidtransToPayment(resMidtrans *coreapi.ChargeResponse) Payment {
	return Payment{
		OrderID:           resMidtrans.OrderID,
		TransactionID:     resMidtrans.TransactionID,
		PaymentMethod:     config.PaymentBankTransferBCA,
		BillNumber:        resMidtrans.VaNumbers[0].VANumber,
		Bank:              resMidtrans.VaNumbers[0].Bank,
		GrossAmount:       resMidtrans.GrossAmount,
		TransactionStatus: resMidtrans.TransactionStatus,
	}
}
