package utils

import "reflect"

// IsEmpty 判断结构体是否为空
func IsEmpty(obj interface{}) bool {
	return reflect.DeepEqual(obj, reflect.Zero(reflect.TypeOf(obj)).Interface())
}
