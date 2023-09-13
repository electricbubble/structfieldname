# StructFieldName

## Install

```shell
go get github.com/electricbubble/structfieldname
```

## Usage

```golang
func ExampleMustLookup_case1() {
	var tmp http.Request
	s := MustLookup(Option{}, &tmp, &tmp.URL)
	fmt.Println(s)

	s = MustLookup(Option{}, &tmp, &tmp.Body)
	fmt.Println(s)

	// Output:
	// URL
	// Body
}

func ExampleMustLookup_case2() {
	type Range struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}
	type Detail struct {
		Name  string `json:"name"`
		Range Range  `json:"range"`

		u struct {
			last uint64
		}
	}

	var tmp Detail

	s := MustLookup(Option{TagKey: "json"}, &tmp, &tmp.Name)
	fmt.Println(s)

	s = MustLookup(Option{Sep: "."}, &tmp, &tmp.Range.Min)
	fmt.Println(s)

	s = MustLookup(Option{TagKey: "json", Sep: "."}, &tmp, &tmp.Range.Max)
	fmt.Println(s)

	s = MustLookup(Option{Sep: "-"}, &tmp, &tmp.u.last)
	fmt.Println(s)

	// Output:
	// name
	// Range.Min
	// range.max
	// u-last
}

```