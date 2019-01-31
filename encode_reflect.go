package proto

import (
	"fmt"
	"reflect"
)

func Encode(data []byte, x interface{}) (out []byte, err error) {
	v := reflect.ValueOf(x)

	switch v.Kind() {
	case reflect.Bool:
		return EncodeBool(data, v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return EncodeSInt64(data, v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return EncodeUInt64(data, v.Uint()), nil
	case reflect.Float32:
		return EncodeFloat32(data, float32(v.Float())), nil
	case reflect.Float64:
		return EncodeFloat64(data, v.Float()), nil
	case reflect.String:
		return EncodeString(data, v.String()), nil
	default:
		// Uintptr, Complex64, Complex128, Array, Chan, Func, Interface, Map, Ptr, Slice, Struct, UnsafePointer
		return nil, fmt.Errorf("unsupported kind %s", v.Kind())
	}
}
