package errors

import (
	"errors"
)

type BusinessError struct {
	error
	Msg    string
	Detail interface{}
	Err    error
}

func (b BusinessError) Error() string {
	return errors.New(b.Msg).Error()
}

func WithMes(msg string) BusinessError {
	return BusinessError{
		Msg:    msg,
		Detail: "",
	}
}

func WithDetail(Key string, detail interface{}) *BusinessError {
	return &BusinessError{
		Msg:    Key,
		Detail: detail,
	}
}

func WithErr(Key string, err error) *BusinessError {
	return &BusinessError{
		Msg:    Key,
		Detail: err,
	}
}
