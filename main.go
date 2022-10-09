package main

import (
	"fmt"
	"time"
	"warehouse/config"
	"warehouse/factory"
	modelCheckout "warehouse/features/checkout/data"
	modelLahan "warehouse/features/lahan/data"
	"warehouse/migration"
	"warehouse/utils/database/mysql"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDBmySql(cfg)

	migration.InitMigrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	gmt, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(gmt)
	s.Every(1).Day().At("00:00").Do(
		func() {
			// Select checkout expired today
			var dataCheckout []modelCheckout.Checkout

			layout_date := "2006-01-02"
			tx_select := db.Where("akhir_sewa = ?", time.Now().Format(layout_date)+" 00:00:00.000").Find(&dataCheckout)

			if tx_select.Error != nil {
				fmt.Println("gagal mencari checkout yang expired")
			} else {
				fmt.Println("berjasil mencari checkout yang expired")
			}

			for _, v := range dataCheckout {
				// Update status checkout
				v.Status = "expired"
				tx_checkoutUpdate := db.Model(&modelCheckout.Checkout{}).Where("id = ?", v.ID).Updates(&v)

				if tx_checkoutUpdate.Error != nil {
					fmt.Println("gagal update status checkout")
				} else {
					fmt.Println("update status checkout berjalan")
				}

				// Update status lahan
				var data modelLahan.Lahan
				tx_selectLahan := db.Where("id = ?", v.LahanID).Preload("Gudang").Find(&data)

				if tx_selectLahan.Error != nil {
					fmt.Println("gagal mencari lahan")
				} else {
					fmt.Println("berhasil mencari lahan")
				}

				data.Status = "tidak disewa"
				tx_update := db.Model(&modelLahan.Lahan{}).Where("id = ?", data.ID).Updates(&data)

				if tx_update.Error != nil {
					fmt.Println("gagal update status lahan")
				} else {
					fmt.Println("update status lahan berjalan")
				}
			}
		},
	)
	s.StartAsync()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
