package datacommon

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

//時間結構
type NullTime struct {
	Time  time.Time
	Valid bool // 是否有值
}

//实现它的赋值方法(注意，这个方属于指针)
func (this *NullTime) Scan(value interface{}) error {
	this.Time, this.Valid = value.(time.Time)
	return nil
}

//实现它的取值方式
func (this NullTime) Value() (driver.Value, error) {
	if !this.Valid {
		return nil, nil
	}
	return this.Time, nil
}

//實現序列成json字串
func (this *NullTime) MarshalJSON() ([]byte, error) {
	if this.Time.IsZero() {
		return []byte("null"), nil
	} else {
		var stamp = fmt.Sprintf("\"%s\"", time.Time(this.Time).Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	}

}

//實現反序列成資料
func (this *NullTime) UnmarshalJSON(data []byte) error {
	var err error
	strDate := string(data)
	strDate = strings.ReplaceAll(strDate, "\"", "")
	this.Time, err = now.Parse(strDate)
	if err == nil {
		this.Valid = true
	}
	return err
}
