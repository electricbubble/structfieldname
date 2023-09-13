package structfieldname

import (
	"testing"
)

func TestLookup(t *testing.T) {
	t.Run("case_s-string", func(t *testing.T) {
		var tmp string
		_, err := Lookup(Option{}, tmp, nil)
		requireEqualError(t, err, ErrMustNonNilStruct.Error())
	})
	t.Run("case_s-&string", func(t *testing.T) {
		var tmp string
		_, err := Lookup(Option{}, &tmp, nil)
		requireEqualError(t, err, ErrMustNonNilStruct.Error())
	})
	t.Run("case_s-*string", func(t *testing.T) {
		var tmp *string
		_, err := Lookup(Option{}, tmp, nil)
		requireEqualError(t, err, ErrMustNonNil.Error())
	})
	t.Run("case_s-&*string", func(t *testing.T) {
		var tmp *string
		_, err := Lookup(Option{}, &tmp, nil)
		requireEqualError(t, err, ErrMustNonNilStruct.Error())
	})

	t.Run("case_s-struct", func(t *testing.T) {
		var tmp tmpStruct1
		_, err := Lookup(Option{}, tmp, nil)
		requireEqualError(t, err, ErrMustNonNilStruct.Error())
	})
	t.Run("case_s-&struct_f-string", func(t *testing.T) {
		var tmp tmpStruct1
		_, err := Lookup(Option{}, &tmp, tmp.BasePath)
		requireEqualError(t, err, ErrMustPointerStructField.Error())
	})
	t.Run("case_s-&struct_f-&string", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.BasePath)
		requireNoError(t, err)
		requireEqual(t, "BasePath", s)

		s, err = Lookup(Option{}, &tmp, &tmp.token)
		requireNoError(t, err)
		requireEqual(t, "token", s)
	})
	t.Run("case_s-&struct_f-&bool", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.Enabled)
		requireNoError(t, err)
		requireEqual(t, "Enabled", s)
	})
	t.Run("case_s-&struct_f-&int", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.NumInt)
		requireNoError(t, err)
		requireEqual(t, "NumInt", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumInt8)
		requireNoError(t, err)
		requireEqual(t, "NumInt8", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumInt16)
		requireNoError(t, err)
		requireEqual(t, "NumInt16", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumInt32)
		requireNoError(t, err)
		requireEqual(t, "NumInt32", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumInt64)
		requireNoError(t, err)
		requireEqual(t, "NumInt64", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumF32)
		requireNoError(t, err)
		requireEqual(t, "NumF32", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumF64)
		requireNoError(t, err)
		requireEqual(t, "NumF64", s)
	})

	t.Run("case_s-&struct_f-&[]int", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.ii)
		requireNoError(t, err)
		requireEqual(t, "ii", s)
	})
	t.Run("case_s-&struct_f-&[]string", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.ss)
		requireNoError(t, err)
		requireEqual(t, "ss", s)
	})
	t.Run("case_s-&struct_f-&chan", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.ch)
		requireNoError(t, err)
		requireEqual(t, "ch", s)
	})
	t.Run("case_s-&struct_f-&func", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp._fn)
		requireNoError(t, err)
		requireEqual(t, "_fn", s)
	})
	t.Run("case_s-&struct_f-&any", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.A)
		requireNoError(t, err)
		requireEqual(t, "A", s)
	})
	t.Run("case_s-&struct_f-&map", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.M)
		requireNoError(t, err)
		requireEqual(t, "M", s)
	})

	t.Run("case_s-&struct_f-&*", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.ptrI32)
		requireNoError(t, err)
		requireEqual(t, "ptrI32", s)

		s, err = Lookup(Option{}, &tmp, &tmp.ptrUI)
		requireNoError(t, err)
		requireEqual(t, "ptrUI", s)
	})

	t.Run("case_s-&struct_f-&f.int", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.NumU.nUint)
		requireNoError(t, err)
		requireEqual(t, "nUint", s)

		s, err = Lookup(Option{}, &tmp, &tmp.NumU.nUint32)
		requireNoError(t, err)
		requireEqual(t, "nUint32", s)
	})

	t.Run("case_s-&struct_f-&f", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.NumU)
		requireNoError(t, err)
		requireEqual(t, "NumU", s)
	})

	t.Run("case_s-&struct_f-&f", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{}, &tmp, &tmp.NumU)
		requireNoError(t, err)
		requireEqual(t, "NumU", s)
	})

	t.Run("case_s-&struct_f-&f.int_opt-key", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{TagKey: "bson"}, &tmp, &tmp.NumU.nUint)
		requireNoError(t, err)
		requireEqual(t, "nUint", s)

		s, err = Lookup(Option{TagKey: "bson"}, &tmp, &tmp.NumU.nUint8)
		requireNoError(t, err)
		requireEqual(t, "nUint8", s)

		s, err = Lookup(Option{TagKey: "bson"}, &tmp, &tmp.NumU.nUint16)
		requireNoError(t, err)
		requireEqual(t, "n_uint16", s)

		s, err = Lookup(Option{TagKey: "bson"}, &tmp, &tmp.NumU.nUint32)
		requireNoError(t, err)
		requireEqual(t, "n_uint32", s)

		s, err = Lookup(Option{TagKey: "bson"}, &tmp, &tmp.NumU.nUint64)
		requireNoError(t, err)
		requireEqual(t, "nUint64", s)
	})

	t.Run("case_s-&struct_f-&f.int_opt-sep", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{Sep: "."}, &tmp, &tmp.other.Port)
		requireNoError(t, err)
		requireEqual(t, "other.Port", s)
	})

	t.Run("case_s-&struct_f-&f.f.f.[]int_opt-key-sep", func(t *testing.T) {
		var tmp tmpStruct1
		s, err := Lookup(Option{TagKey: "bson", Sep: "."}, &tmp, &tmp.SSS.INT.LAST.values)
		requireNoError(t, err)
		requireEqual(t, "SSS.int.last.values", s)

		s, err = Lookup(Option{TagKey: "bson", Sep: "_"}, &tmp, &tmp.SSS.INT.LAST.values)
		requireNoError(t, err)
		requireEqual(t, "SSS_int_last_values", s)

		s, err = Lookup(Option{TagKey: "bson", Sep: "."}, &tmp, &tmp.SSS.INT.LAST)
		requireNoError(t, err)
		requireEqual(t, "SSS.int.last", s)
	})
}

type tmpStruct1 struct {
	Enabled  bool   `bson:"enabled,omitempty"`
	BasePath string `bson:"base_path,omitempty"`

	NumInt   int   `bson:"-"`
	NumInt8  int8  `bson:"-"`
	NumInt16 int16 `bson:"-"`
	NumInt32 int32 `bson:"-"`
	NumInt64 int64 `bson:"-"`
	NumF32   float32
	NumF64   float64

	NumU struct {
		nUint   uint `bson:""`
		nUint8  uint `bson:"	"`
		nUint16 uint `bson:"n_uint16"`
		nUint32 uint `bson:"n_uint32,omitempty"`
		nUint64 uint `bson:"-"`
	} `bson:"num_u"`

	other struct {
		Port int `bson:"port,omitempty"`
	} `bson:"other,omitempty"`
	token    string   `bson:"token,omitempty"`
	internal internal `bson:"internal,omitempty"`

	ii []int
	ss []string `bson:"ss,omitempty"`

	Public Public `bson:"public,omitempty"`
	SSS    struct {
		INT struct {
			LAST struct {
				values []int `bson:"values,omitempty"`
			} `bson:"last,omitempty"`
		} `bson:"int,omitempty"`

		keys  []string `bson:"keys,omitempty"`
		total []any    `bson:"total,omitempty"`
	} `bson:"SSS,omitempty"`

	ch  chan struct{} `bson:"ch,omitempty"`
	_fn func()
	A   any
	M   map[string]any

	ptrI32 *int32 `bson:"ptr_int32,omitempty"`
	ptrUI  *uint  `bson:"ptr_uint,omitempty"`
}

type internal struct {
	secret string `bson:"secret,omitempty"`
}

type Public struct {
	Name string `bson:"name,omitempty"`
}
