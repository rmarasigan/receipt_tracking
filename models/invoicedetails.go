package models

import (
	"fmt"

	"github.com/uadmin/uadmin"
)

type InvoiceDetail struct {
	uadmin.Model
	Invoice     Invoice `uadmin:"list_exclude"`
	InvoiceID   uint
	ItemName    string  `uadmin:"required"`
	Quantity    int     `uadmin:"required"`
	UnitPrice   float64 `uadmin:"required"`
	TotalAmount float64 `uadmin:"read_only"`
}

func (i *InvoiceDetail) String() string {
	invoice := Invoice{}
	uadmin.Get(&invoice, "id = ?", i.InvoiceID)
	invoiceNumber := invoice.InvoiceNumber

	return fmt.Sprintf("%v", invoiceNumber)
}

func (i *InvoiceDetail) Save() {
	// Automatically compute for the total amount
	i.TotalAmount = (float64(i.Quantity) * i.UnitPrice)

	// Update Invoice Total Amount on each save of Invoice Detail
	invoice := Invoice{}
	uadmin.Get(&invoice, "id = ?", i.InvoiceID)

	if invoice.ID == i.InvoiceID && true {
		totalAmount := invoice.TotalAmount + i.TotalAmount
		uadmin.Update(&invoice, "total_amount", totalAmount, "id = ?", i.InvoiceID)

		uadmin.Save(i)
	}
	uadmin.Save(i)
}
