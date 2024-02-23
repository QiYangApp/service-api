package response

type StateEnum string
type TypeEnum string

const (
	Success StateEnum = "success"
	Error   StateEnum = "error"
	Warn    StateEnum = "warn"
	Fail    StateEnum = "fail"
)

const (
	IMAGE     TypeEnum = "IMAGE"
	FILE      TypeEnum = "FILE"
	JSON      TypeEnum = "JSON"
	String    TypeEnum = "String"
	AsciiJson TypeEnum = "AsciiJSON"
	ProtoBuf  TypeEnum = "ProtoBuf"
	PureJson  TypeEnum = "PureJson"
	XML       TypeEnum = "XML"
	YAML      TypeEnum = "YAML"
	DATA      TypeEnum = "DATA"
)

type ResponseMethods[T interface{}] interface {
	// Single 单个返回
	Single(data T) Response
	// Pagination 分页
	Pagination(list []T) Response
	// Multiple 多个返回
	Multiple(data []T) Response
}
