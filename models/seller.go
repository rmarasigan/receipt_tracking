package models

import "github.com/uadmin/uadmin"

type Seller struct {
	uadmin.Model
	Code    string `uadmin:"required"`
	Name    string `uadmin:"required"`
	Address string
}
