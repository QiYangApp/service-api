package auth

import (
	authtype "ent/types/auth"
	"reflect"
)

var registeredConfigs = map[authtype.Type]func() authtype.ConfigType[any]{}

// RegisterTypeConfig register a config for a provided type
func RegisterTypeConfig(typ authtype.Type, exemplar authtype.ConfigType[any]) {
	if reflect.TypeOf(exemplar).Kind() == reflect.Ptr {
		// Pointer:
		registeredConfigs[typ] = func() authtype.ConfigType[any] {
			return reflect.New(reflect.ValueOf(exemplar).Elem().Type()).Interface().(authtype.ConfigType[any])
		}
		return
	}

	// Not a Pointer
	registeredConfigs[typ] = func() authtype.ConfigType[any] {
		return reflect.New(reflect.TypeOf(exemplar)).Elem().Interface().(authtype.ConfigType[any])
	}
}
