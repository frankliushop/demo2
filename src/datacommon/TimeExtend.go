package datacommon

import (
	"database/sql/driver"
	"fmt"

	//"os"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

func init() {
	//os.Setenv("TZ", "Asia/Taipei")
	fmt.Println("load time zone")
	timelocal, _ := time.LoadLocation("Asia/Taipei")
	time.Local = timelocal
	//time.LoadLocation("Asia/Taipei")
}

type NullTime struct {
	Time  time.Time
	Valid bool // 是否有值
}

//实现它的赋值方法(注意，这个方属于指针)
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

//实现它的取值方式
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (this *NullTime) MarshalJSON() ([]byte, error) {
	if this.Time.IsZero() {
		return nil, nil
	} else {
		var stamp = fmt.Sprintf("\"%s\"", time.Time(this.Time).Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	}

}

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
