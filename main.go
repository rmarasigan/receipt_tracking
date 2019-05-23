package main

import (
	"github.com/rmarasigan/receipt_tracking/models"
	"github.com/uadmin/uadmin"
)

func main() {

	// Database Configuration
	dbSettings := uadmin.DBSettings{}
	dbSettings.Name = "receipt_tracker"
	dbSettings.Type = "sqlite"
	uadmin.Database = &dbSettings

	uadmin.SiteName = "Receipt Tracker"
	uadmin.RootURL = "/receipt-tracker/"

	uadmin.Register(
		models.Seller{},
		models.Customer{},
		models.Invoice{},
		models.InvoiceDetail{},
	)

	uadmin.Port = 2563
	uadmin.StartServer()
}
