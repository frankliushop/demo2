package datacommon

import (
	"fmt"
	"time"
)

type NullDateExtend struct {
	NullTimeExtend
}

func (this *NullDateExtend) MarshalJSON() ([]byte, error) {
	if this.Time.IsZero() {
		return []byte("null"), nil
	} else {
		var stamp = fmt.Sprintf("\"%s\"", time.Time(this.Time.Time).Format("2006-01-02"))
		return []byte(stamp), nil
	}

}
