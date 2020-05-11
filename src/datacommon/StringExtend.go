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
func (ns *NullString) Scan(value interface{}) error {
	ns.Valid = false
	if value != nil {
		var strData string
		byteDatas, _ := value.([]byte)
		strData = string(byteDatas)
		ns.Data = &strData
		ns.Valid = true
	}
	return nil
}

//实现它的取值方式
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return *ns.Data, nil
}

//實現序列成json字串
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	} else {
		return []byte(*ns.Data), nil
	}

}

//實現反序列成資料
func (ns *NullString) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData != "null" {
		ns.Data = &strData
	}
	return nil
}
