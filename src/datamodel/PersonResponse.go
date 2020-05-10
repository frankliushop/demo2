package datamodel

import (
	"myproject1/datacommon"
)

type PersonResponse struct {
	ID          int32          `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	MobilePhone string    `json:"mobilePhone"`
	Address     string    `json:"address"`
	Birthday    datacommon.NullTime `json:"birthday"`
}
