package data

import (
	"fmt"
	"warehouse/features/checkout"

	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type checkoutData struct {
	db *gorm.DB
}

func New(db *gorm.DB) checkout.DataInterface {
	return &checkoutData{
		db: db,
	}
}

func (repo *checkoutData) AddCheckoutByFav(data checkout.Core) (int, error) {

	dataModel := fromCore(data)
	fmt.Println("Query :", data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *checkoutData) CreateDataPayment(reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	payment, errPayment := coreapi.ChargeTransaction(&reqPay)
	if errPayment != nil {
		return nil, errPayment.RawError
	}
	return payment, nil
}

func (repo *checkoutData) SelectPayment(orderID string) (checkout.Core, error) {
	payment := Checkout{}

	findData := repo.db.Order("id desc").Where("order_id = ?", orderID).Find(&payment)
	if findData.Error != nil {
		return checkout.Core{}, findData.Error
	}
	result := payment.toCoreMidtrans()
	return result, nil
}

func (repo *checkoutData) PaymentDataWebHook(data checkout.Core) error {
	Model := fromCore(data)
	if data.Status == "paid" {
		errUpdateStatus := repo.db.Where("order_id = ?", data.OrderID).Model(&Model).Update("status", data.Status)
		if errUpdateStatus.Error != nil {
			return errUpdateStatus.Error
		}
		return nil
	} else {
		errUpdate := repo.db.Where("order_id = ?", data.OrderID).Model(&Model).Updates(Checkout{
			Status:           data.Status,
			MetodePembayaran: data.MetodePembayaran,
			TransactionID:    data.TransactionID,
		})
		if errUpdate.Error != nil {
			return errUpdate.Error
		}
		return nil
	}
}
