package enums

import (
	"fmt"
	"strings"
)

// Uinter defines required (read-only) methods
// for all Uint-derived types interface.
type Uinter interface {
	Touint() Uint
	// Marshal returns next Uinter, string repr of current Uinter and an error if any
	Marshal() (Uinter, string, error)
	// MarshalJSON() ([]byte, error)
}

// Upointer dictated methods:

// Once there're no renamed parameters, these Unmarshal should do.
// func (r *UintDF) Unmarshal(s string) error { return UnmarshalStringFunc(r.UnmarshalJSON)(s) }
// func (r *UintPS) Unmarshal(s string) error { return UnmarshalStringFunc(r.UnmarshalJSON)(s) }

// Unmarshal for UintDF. Knows renamed parameter.
func (r *UintDF) Unmarshal(data string, negate *bool) error {
	issize := data == "size" || data == "dfsize"
	if issize {
		data = "total"
	}
	isexnondeffs := data == "fs" && !*negate
	if isexnondeffs { // "df=fs" became "df=-fs"
		*negate = !*negate
	}
	return UnmarshalMaybe(isexnondeffs || issize, r.UnmarshalJSON, data)
}

// Unmarshal for UintPS. Knows renamed parameter.
func (r *UintPS) Unmarshal(data string, negate *bool) error {
	issize := data == "size" || data == "pssize"
	if issize {
		data = "virt"
	}
	/* isuser := data == "user"
	if isuser {
		data = "uid"
	} // */
	isexnondefps := data == "pid" && !*negate
	if isexnondefps { // "ps=pid" became "ps=-pid"
		*negate = !*negate
	}
	return UnmarshalMaybe(isexnondefps || issize /* || isuser */, r.UnmarshalJSON, data)
}

// Uinter dictated methods:

func (r UintDF) Touint() Uint                     { return Uint(r) }
func (r UintDF) Marshal() (Uinter, string, error) { return MarshalStringFunc(r+1, r.MarshalJSON)() }

func (r UintDFT) Touint() Uint                     { return Uint(r) }
func (r UintDFT) Marshal() (Uinter, string, error) { return MarshalStringFunc(r+1, r.MarshalJSON)() }
func (r *UintDFT) Unmarshal(data string, _ *bool) error {
	return UnmarshalMaybe(false, r.UnmarshalJSON, data)
}

func (r UintIFT) Touint() Uint                     { return Uint(r) }
func (r UintIFT) Marshal() (Uinter, string, error) { return MarshalStringFunc(r+1, r.MarshalJSON)() }
func (r *UintIFT) Unmarshal(data string, _ *bool) error {
	return UnmarshalMaybe(false, r.UnmarshalJSON, data)
}

func (r UintPS) Touint() Uint                     { return Uint(r) }
func (r UintPS) Marshal() (Uinter, string, error) { return MarshalStringFunc(r+1, r.MarshalJSON)() }

// Helpers:

func UnmarshalMaybe(rename bool, unmarshal BytesUnmarshal, data string) error {
	if err := UnmarshalStringFunc(unmarshal)(data); err != nil || !rename {
		return err
	}
	return RenamedConstError("")
}

func UnmarshalStringFunc(unmarshaler BytesUnmarshal) func(string) error {
	return func(data string) error {
		return unmarshaler([]byte(fmt.Sprintf("%q", strings.ToUpper(data))))
	}
}

func MarshalStringFunc(next Uinter, marshaler BytesEnmarshal) func() (Uinter, string, error) {
	return func() (Uinter, string, error) {
		b, err := marshaler()
		if err != nil {
			return next, "", err
		}
		if l := len(b); l > 2 && b[0] == '"' && b[l-1] == '"' {
			b = b[1 : l-1]
		}
		s := strings.ToLower(string(b))
		return next, s, nil
	}
}

type BytesEnmarshal func() ([]byte, error)
type BytesUnmarshal func([]byte) error

// RenamedConstError denotes an error.
type RenamedConstError string

func (rc RenamedConstError) Error() string { return string(rc) }
