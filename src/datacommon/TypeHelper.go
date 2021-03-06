package datacommon

import (
	"gopkg.in/guregu/null.v3"
)

//字串是否爲空字串或nil
func IsNullOrEmpty(data null.String) bool {
	if data.Valid && data.String != "" {
		return false
	}
	return true
}

//檢查數字是否為nil或為0
func IsNullOrZero(data null.Int) bool {
	if data.Valid && data.Int64 > 0 {
		return false
	}
	return true
}
