package personmodel

import (
	"myproject1/datacommon"

	"gopkg.in/guregu/null.v3"
)

type PersonResponse struct {
	ID          null.Int                  `json:"id"`
	Name        null.String               `json:"name"`
	Phone       null.String               `json:"phone"`
	MobilePhone null.String               `json:"mobilePhone"`
	Address     null.String               `json:"address"`
	Birthday    datacommon.NullDateExtend `json:"birthday"`
}
