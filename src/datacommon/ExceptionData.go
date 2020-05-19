package datacommon

import (
	"fmt"
)


//例外處理
type ExceptionData struct {
	error
	//錯誤碼
	ErrorCode    int16
	//錯誤訊息
	ErrorMessage string	
}

func (this *ExceptionData) String() string{
	return fmt.Sprintf("errorCode : %d , errorMessage : %s",this.ErrorCode,this.ErrorMessage)
}
