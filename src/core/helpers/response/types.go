package response

type ResponseStateEnum string

const (
	Success ResponseStateEnum = "success"
	Error   ResponseStateEnum = "error"
	Warn    ResponseStateEnum = "warn"
	Fail    ResponseStateEnum = "fail"
)

type ResponseMethods[T interface{}] interface {
	// Single 单个返回
	Single(data T) Response[T]
	// Pagination 分页
	Pagination(list []T) Response[T]
	// Multiple 多个返回
	Multiple(data []T) Response[T]
}
