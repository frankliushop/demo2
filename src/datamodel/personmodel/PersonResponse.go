package personmodel

import (
	"myproject1/datacommon"
)

type PersonResponse struct {
	ID          int32                 `json:"id"`
	Name        datacommon.NullString `json:"name"`
	Phone       datacommon.NullString `json:"phone"`
	MobilePhone datacommon.NullString `json:"mobilePhone"`
	Address     datacommon.NullString `json:"address"`
	Birthday    datacommon.NullTime   `json:"birthday"`
}
