package datacommon

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"gopkg.in/guregu/null.v4"
)

type NullTimeExtend struct {
	null.Time
}

func (this *NullTimeExtend) UnmarshalJSON(data []byte) error {
	var err error
	strDate := string(data)
	strDate = strings.ReplaceAll(strDate, "\"", "")
	this.Time.Time, err = now.Parse(strDate)
	if err == nil {
		this.Valid = true
	}
	return err
}

func (this *NullTimeExtend) MarshalJSON() ([]byte, error) {
	if this.Time.IsZero() {
		return []byte("null"), nil
	} else {
		var stamp = fmt.Sprintf("\"%s\"", time.Time(this.Time.Time).Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	}

}
