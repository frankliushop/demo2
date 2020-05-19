package datacommon

import (
	"database/sql/driver"
)

//字串結構
type NullString struct {
	Data  *string
	Valid bool // 是否有值
}

//实现它的赋值方法(注意，这个方属于指针)
func (this *NullString) Scan(value interface{}) error {
	this.Valid = false
	if value != nil {
		var strData string
		byteDatas, _ := value.([]byte)
		strData = string(byteDatas)
		this.Data = &strData
		this.Valid = true
	}
	return nil
}

//实现它的取值方式
func (this NullString) Value() (driver.Value, error) {
	if !this.Valid {
		return nil, nil
	}
	return *this.Data, nil
}

//實現序列成json字串
func (this *NullString) MarshalJSON() ([]byte, error) {
	if !this.Valid {
		return []byte("null"), nil
	} else {
		return []byte(*this.Data), nil
	}

}

//實現反序列成資料
func (this *NullString) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData != "null" {
		this.Data = &strData
	}
	return nil
}
