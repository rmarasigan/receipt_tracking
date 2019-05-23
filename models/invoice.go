package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

type Invoice struct {
	uadmin.Model
	InvoiceNumber string `uadmin:"required"`
	Seller        Seller
	SellerID      uint
	Customer      Customer
	CustomerID    uint
	Date          time.Time `uadmin:"required"`
	TotalAmount   string    `uadmin:"required"`
}

func (i Invoice) String() string {
	seller := Seller{}
	uadmin.Get(&seller, "id = ?", i.SellerID)
	sellerName := seller.Code

	return fmt.Sprintf("%v | %v", i.InvoiceNumber, sellerName)
}
