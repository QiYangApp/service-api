package inject

// DiService  Apis Project global API cache
var DI map[string]*MethodInfo

// Annotations the annotation of Api method
type Annotations map[string]string

// Global di cache
var diCache []inject

// Annotations of each API
var annotationCache map[string]Annotations

// Register controllers
func Register(inject ...inject) {
	diCache = append(diCache, inject...)
}
