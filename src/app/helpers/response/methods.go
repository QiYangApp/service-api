package response

import "time"

func NewResponse[T interface{}]() *Response[T] {
	return &Response[T]{
		Timestamp: time.Now(),
	}
}

func Single[T interface{}](data T) *Response[T] {
	return NewResponse[T]().Single(data)
}

func SingleSuccess[T interface{}](data T) *Response[T] {
	return NewResponse[T]().Single(data).SetCode(200).SetMessage("SUCCESS")
}

func SingleCustom[T interface{}](data T, code int, message string) *Response[T] {
	return NewResponse[T]().Single(data).SetCode(code).SetMessage(message)
}
