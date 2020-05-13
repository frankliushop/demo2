package datacommon

import (
	"gopkg.in/guregu/null.v4"
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
	if data.Valid && !data.IsZero() {
		return false
	}
	return true
}
