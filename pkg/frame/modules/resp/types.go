package resp

type StateEnum string
type TypeEnum string

const (
	SuccessState StateEnum = "success"
	ErrorState   StateEnum = "error"
	WarnState    StateEnum = "warn"
	FailState    StateEnum = "fail"
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
