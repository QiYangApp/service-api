package response

type ResponseStateEnum string
type ResponseTypeEnum string

const (
	Success ResponseStateEnum = "success"
	Error   ResponseStateEnum = "error"
	Warn    ResponseStateEnum = "warn"
	Fail    ResponseStateEnum = "fail"
)

const (
	IMAGE     ResponseTypeEnum = "IMAGE"
	FILE      ResponseTypeEnum = "FILE"
	JSON      ResponseTypeEnum = "JSON"
	String    ResponseTypeEnum = "String"
	AsciiJson ResponseTypeEnum = "AsciiJSON"
	ProtoBuf  ResponseTypeEnum = "ProtoBuf"
	PureJson  ResponseTypeEnum = "PureJson"
	XML       ResponseTypeEnum = "XML"
	YAML      ResponseTypeEnum = "YAML"
	DATA      ResponseTypeEnum = "DATA"
)

type ResponseMethods[T interface{}] interface {
	// Single 单个返回
	Single(data T) Response[T]
	// Pagination 分页
	Pagination(list []T) Response[T]
	// Multiple 多个返回
	Multiple(data []T) Response[T]
}
