package errors

import (
	"errors"
)

type BusinessError struct {
	Msg    string
	Detail interface{}
	Err    error
}

func (e BusinessError) Error() string {
	return errors.New(e.Msg).Error()
}

func WithMes(Key string) *BusinessError {
	return &BusinessError{
		Msg:    Key,
		Detail: "",
		Err:    nil,
	}
}

func WithDetail(Key string, detail interface{}) *BusinessError {
	return &BusinessError{
		Msg:    Key,
		Detail: detail,
		Err:    nil,
	}
}

func WithErr(Key string, err error) *BusinessError {
	return &BusinessError{
		Msg:    Key,
		Detail: "",
		Err:    err,
	}
}

func WithDetailAndErr(Key string, detail interface{}, err error) BusinessError {
	return BusinessError{
		Msg:    Key,
		Detail: detail,
		Err:    err,
	}
}
