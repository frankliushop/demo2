package datamodel

type ResultResponse struct {
	Result  bool	`json:"result"`
	Code    int16	`json:"code"`
	Message string	`json:"message"`
	DataResult    interface{}	`json:"dataResult"`
}
