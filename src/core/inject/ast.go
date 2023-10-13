package inject

// Apis Project global API cache
var Apis map[string]*MethodInfo

// MethodInfo Api method info
type MethodInfo struct {
	Method      string // API method。such as: POST、GET、DELETE、PUT、OPTIONS、PATCH、HEAD
	ApiPath     string // API path
	PackPath    string // API path
	PackName    string
	Annotations map[string]string // Annotations of the method
}