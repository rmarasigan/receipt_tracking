package models

import "github.com/uadmin/uadmin"

type Customer struct {
	uadmin.Model
	Name string `uadmin:"required"`
}
