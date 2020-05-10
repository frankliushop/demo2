package datamodel

type ResultResponse struct {
	Result bool
	Code int16
	Message string
	Data   interface{}
}
