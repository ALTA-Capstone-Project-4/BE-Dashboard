package usecase

import (
	"errors"
	"warehouse/features/checkout"
	"warehouse/features/lahan"

	"github.com/midtrans/midtrans-go/coreapi"
)

type checkoutUsecase struct {
	checkoutData checkout.DataInterface
	lahanData    lahan.DataInterface
}

func New(data checkout.DataInterface, dataLahan lahan.DataInterface) checkout.UsecaseInterface {
	return &checkoutUsecase{
		checkoutData: data,
		lahanData:    dataLahan,
	}
}

func (usecase *checkoutUsecase) PostCheckoutByFav(data checkout.Core) (int, error) {

	corelahan, errlahan := usecase.lahanData.SelectDetailLahan(data.LahanID, "penitip")
	if errlahan != nil {
		return -1, errlahan
	}
	data.TotalHarga = corelahan.Harga

	row, err := usecase.checkoutData.AddCheckoutByFav(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *checkoutUsecase) GetHargaLahan(lahan_id int, role string) (int, error) {

	corelahan, errlahanHarga := usecase.lahanData.SelectDetailLahan(lahan_id, role)
	if errlahanHarga != nil {
		return 0, errlahanHarga
	}

	return corelahan.Harga, nil
}

func (usecase *checkoutUsecase) CreatePaymentBankTransfer(reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	createPay, errCreatePay := usecase.checkoutData.CreateDataPayment(reqPay)
	if errCreatePay != nil {
		return nil, errors.New("failed get response payment")
	}

	return createPay, nil
}

func (usecase *checkoutUsecase) PaymentWebHook(orderID, status string) error {
	payment, errPayment := usecase.checkoutData.SelectPayment(orderID)
	if errPayment != nil {
		return errors.New("failed to get data join")
	}

	if status == "settlement" {
		payment.Status = "paid"
	}
	if status == "cancel" || status == "deny" || status == "expire" {
		payment.Status = "failed"
		payment.MetodePembayaran = ""
		payment.TransactionID = ""
	}

	result := usecase.checkoutData.PaymentDataWebHook(payment)
	if result != nil {
		return result
	}
	return nil
}
