package structfieldname

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrMustNonNil             = errors.New("must be a non-nil pointer")
	ErrMustNonNilStruct       = errors.New("must be a non-nil struct pointer")
	ErrMustPointerStructField = errors.New("must be a pointer to a struct field")
	ErrUnreachable            = errors.New("unreachable")
)

type Option struct {
	TagKey string
	Sep    string

	prefix string
}

func MustLookup(opt Option, ptrStruct, ptrField any) string {
	s, err := Lookup(opt, ptrStruct, ptrField)
	if err != nil {
		panic(err)
	}
	return s
}

func Lookup(opt Option, ptrStruct, ptrField any) (string, error) {
	rvStruct := reflect.ValueOf(ptrStruct)
	if rvStruct.Kind() != reflect.Pointer {
		return "", ErrMustNonNilStruct
	}
	if rvStruct.IsNil() {
		return "", ErrMustNonNil
	}
	rvStruct = rvStruct.Elem()
	if rvStruct.Kind() != reflect.Struct {
		return "", ErrMustNonNilStruct
	}

	rvField := reflect.ValueOf(ptrField)
	if rvField.Kind() != reflect.Pointer {
		return "", ErrMustPointerStructField
	}

	return lookup(&opt, rvStruct, rvField)
}

func lookup(opt *Option, rvStruct, rvField reflect.Value) (string, error) {
	dstPtr := rvField.Pointer()
	dstTyp := rvField.Elem().Type()

	for i := 0; i < rvStruct.NumField(); i++ {
		vField := rvStruct.Field(i)
		switch vField.Kind() {
		case reflect.String, reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Array, reflect.Slice,
			reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer:
			if uintptr(vField.Addr().UnsafePointer()) == dstPtr {
				return opt.yield(rvStruct.Type().Field(i)), nil
			}
			continue
		case reflect.Struct:
			stf := rvStruct.Type().Field(i)
			if uintptr(vField.Addr().UnsafePointer()) == dstPtr {
				if stf.Type == dstTyp {
					return opt.yield(stf), nil
				}
			}
			s, err := lookup(opt, vField, rvField)
			if errors.Is(err, ErrUnreachable) {
				continue
			}
			if err != nil {
				return "", err
			}
			return opt.join(stf, s), nil
		}
	}

	return "", ErrUnreachable
}

func (opt *Option) yield(f reflect.StructField) string {
	switch {
	case opt.TagKey == "" && opt.Sep == "":
		return f.Name
	case opt.TagKey == "" && opt.Sep != "" && opt.prefix == "":
		return f.Name
	case opt.TagKey == "" && opt.Sep != "" && opt.prefix != "":
		return opt.prefix + opt.Sep + f.Name
	case opt.TagKey != "" && opt.Sep == "":
		return getStructFieldTagValue(f, opt.TagKey)
	case opt.TagKey != "" && opt.Sep != "" && opt.prefix == "":
		return getStructFieldTagValue(f, opt.TagKey)
	case opt.TagKey != "" && opt.Sep != "" && opt.prefix != "":
		return opt.prefix + opt.Sep + getStructFieldTagValue(f, opt.TagKey)
	default:
		panic(ErrUnreachable)
	}
}

func (opt *Option) join(f reflect.StructField, s string) string {
	opt.prefix = getStructFieldTagValue(f, opt.TagKey)

	switch {
	case opt.Sep == "":
		return s
	case opt.Sep != "" && opt.prefix == "":
		return s
	case opt.Sep != "" && opt.prefix != "":
		return opt.prefix + opt.Sep + s
	default:
		panic(ErrUnreachable)
	}
}

func getStructFieldTagValue(f reflect.StructField, key string) string {
	tagValue := strings.TrimSpace(f.Tag.Get(key))
	if tagValue == "" || tagValue == "-" {
		return f.Name
	}
	parts := strings.SplitN(tagValue, ",", 2)
	if parts[0] == "" {
		return f.Name
	}
	return parts[0]
}
