package proto

import (
	"fmt"
	"strconv"
	"reflect"
)

func Encode(x interface{}) (out []byte, err error) {
	v := reflect.ValueOf(x)

	size, err := reflectSize(v)
	if err != nil {
		return nil, err
	}

	out = make([]byte, 0, size)
	return reflectEncode(out, v)
}

func reflectSize(v reflect.Value) (size int, err error) {
	switch v.Kind() {
	case reflect.Bool:
		return SizeBool(v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return SizeSInt64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return SizeUInt64(v.Uint()), nil
	case reflect.Float32:
		return SizeFloat32(float32(v.Float())), nil
	case reflect.Float64:
		return SizeFloat64(v.Float()), nil
	case reflect.String:
		return SizeString(v.String()), nil
	case reflect.Struct:
		vt := v.Type()

		for i := 0; i < vt.NumField(); i += 1 {
			f := vt.Field(i)
			fv := v.Field(i)
			fk := f.Type.Kind()

			str, ok := f.Tag.Lookup("proto")
			if !ok {
				return 0, fmt.Errorf("missing tag for field %s", f.Name)
			}

			id, err := strconv.Atoi(str)
			if err != nil {
				return 0, fmt.Errorf("unable to parse tag for field %s", f.Name)
			}

			switch fk {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				size += SizeTag(id)
			case reflect.Float64:
				size += SizeTag(id)
			case reflect.String, reflect.Struct:
				size += SizeTag(id)

				// TODO this is called twice
				s, err := reflectSize(fv)
				if err != nil {
					return 0, err
				}

				size += SizeInt(s)
			case reflect.Float32:
				size += SizeTag(id)
			default:
				// Uintptr, Complex64, Complex128, Array, Chan, Func, Interface, Map, Ptr, Slice, UnsafePointer
				return 0, fmt.Errorf("unsupported type of kind %s", fk)
			}

			s, err := reflectSize(fv)
			if err != nil {
				return 0, err
			}

			size += s
		}

		return size, nil
	default:
		// Uintptr, Complex64, Complex128, Array, Chan, Func, Interface, Map, Ptr, Slice, UnsafePointer
		return 0, fmt.Errorf("unsupported type of kind %s", v.Kind())
	}
}

func reflectEncode(out []byte, v reflect.Value) ([]byte, error) {
	switch v.Kind() {
	case reflect.Bool:
		return EncodeBool(out, v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return EncodeSInt64(out, v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return EncodeUInt64(out, v.Uint()), nil
	case reflect.Float32:
		return EncodeFloat32(out, float32(v.Float())), nil
	case reflect.Float64:
		return EncodeFloat64(out, v.Float()), nil
	case reflect.String:
		return EncodeString(out, v.String()), nil
	case reflect.Struct:
		vt := v.Type()

		for i := 0; i < vt.NumField(); i += 1 {
			f := vt.Field(i)
			fv := v.Field(i)
			fk := f.Type.Kind()

			str, ok := f.Tag.Lookup("proto")
			if !ok {
				return nil, fmt.Errorf("missing tag for field %s", f.Name)
			}

			id, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("unable to parse tag for field %s", f.Name)
			}

			switch fk {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				out = EncodeTag(out, id, WIRETYPE_VARINT)
			case reflect.Float64:
				out = EncodeTag(out, id, WIRETYPE_64BIT)
			case reflect.String, reflect.Struct:
				out = EncodeTag(out, id, WIRETYPE_LENGTH)

				s, err := reflectSize(fv)
				if err != nil {
					return nil, err
				}

				out = EncodeInt(out, s)
			case reflect.Float32:
				out = EncodeTag(out, id, WIRETYPE_32BIT)
			default:
				// Uintptr, Complex64, Complex128, Array, Chan, Func, Interface, Map, Ptr, Slice, UnsafePointer
				return nil, fmt.Errorf("unsupported type of kind %s", fk)
			}

			out, err = reflectEncode(out, fv)
			if err != nil {
				return nil, err
			}
		}

		return out, nil
	default:
		// Uintptr, Complex64, Complex128, Array, Chan, Func, Interface, Map, Ptr, Slice, UnsafePointer
		return nil, fmt.Errorf("unsupported type of kind %s", v.Kind())
	}
}
