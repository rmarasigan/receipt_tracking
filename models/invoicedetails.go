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
	TotalAmount float64 `uadmin:"required"`
}

func (d *InvoiceDetail) String() string {
	invoice := Invoice{}
	uadmin.Get(&invoice, "id = ?", d.InvoiceID)
	invoiceNumber := invoice.InvoiceNumber

	return fmt.Sprintf("%v", invoiceNumber)
}
