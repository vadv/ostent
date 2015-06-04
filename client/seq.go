package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"sort"
	"strings"

	"github.com/ostrost/ostent/client/enums"
)

// Number is an enums.Uint with sign.
type Number struct {
	enums.Uint
	Negative bool
}

// Upointer defines required (incl. pointer-) methods
// for all enums.Uint-derived types interface.
type Upointer interface {
	enums.Uinter
	Unmarshal(string, *bool) error
	// UnmarshalJSON([]byte) error
}

// DropLink has drop{down,up} link attributes.
type DropLink struct {
	CLASSNAME  string // required for jsx
	AlignClass string
	Text       string `json:"-"` // static
	Href       string
	Class      string
	CaretClass string
}

// EncodeUint returns enums.uinter applied DropLink. .AlignClass is not filled.
func (ep EnumParam) EncodeUint(pname string, uinter enums.Uinter) DropLink {
	base := ep.Params.ValuesCopy()
	text, cur := ep.SetBase(base, pname, uinter)
	dl := DropLink{Text: text, Class: "state"}
	if cur != nil {
		dl.CaretClass = "caret"
		dl.Class += " current"
		if *cur {
			dl.Class += " dropup"
		}
	}
	dl.Href = "?" + ValuesEncode(base)
	return dl
}

// SetBase modifies the base.
func (ep EnumParam) SetBase(base url.Values, pname string, uinter enums.Uinter) (string, *bool) {
	this := uinter.Touint()
	_, low, err := uinter.Marshal()
	if err != nil { // ignoring the error
		return "", nil
	}

	text := ep.EnumDecodec.Text(strings.ToUpper(low))
	ddef := ep.EnumDecodec.Default.Uint
	dnum := ep.EnumDecoded.Number

	// Default ordering is desc (values are numeric most of the time).
	// Alpha values ordering: asc.
	desc := !ep.IsAlpha(this)
	if dnum.Negative {
		desc = !desc
	}
	var ret *bool
	if this == dnum.Uint {
		ret = new(bool)
		*ret = !desc
	}
	// for default, opposite of having a parameter is it's absence.
	if this == ddef && ep.EnumDecoded.Specified {
		base.Del(pname)
		return text, ret
	}
	if this == dnum.Uint && !dnum.Negative {
		low = "-" + low
	}
	base.Set(pname, low)
	return text, ret
}

type EnumDecodec struct {
	Default Number
	Alphas  []enums.Uint
	Unew    func() (string, Upointer) `json:"-"` // func cannot be marshaled
	Text    func(string) string       `json:"-"` // func cannot be marshaled
	Pname   string
}

func (ep EnumParam) IsAlpha(p enums.Uint) bool {
	for _, u := range ep.EnumDecodec.Alphas {
		if u == p {
			return true
		}
	}
	return false
}

func (ed EnumDecodec) DefaultParam(params *Params) EnumParam {
	return EnumParam{
		EnumDecodec: ed,
		Params:      params,
	}
}

// TextFunc constructs string replacement func.
// ab defines exact mapping, miss-case: fs funcs chain-apply.
func TextFunc(ab map[string]string, fs ...func(string) string) func(string) string {
	return func(s string) string {
		if n, ok := ab[s]; ok {
			return n
		}
		for _, f := range fs {
			s = f(s)
		}
		return s
	}
}

var EnumDecodecs = map[string]EnumDecodec{
	"ps": {
		Default: Number{Uint: enums.Uint(enums.PID)},
		Alphas:  []enums.Uint{enums.Uint(enums.NAME), enums.Uint(enums.USER)},
		Unew:    func() (string, Upointer) { return "ps", new(enums.UintPS) },
		Text:    TextFunc(map[string]string{"PRI": "PR", "NICE": "NI", "NAME": "COMMAND"}, strings.ToUpper),
	},
	"df": {
		Default: Number{Uint: enums.Uint(enums.FS)},
		Alphas:  []enums.Uint{enums.Uint(enums.FS), enums.Uint(enums.MP)},
		Unew:    func() (string, Upointer) { return "df", new(enums.UintDF) },
		Text:    TextFunc(map[string]string{"FS": "Device", "MP": "Mounted"}, strings.ToLower, strings.Title),
	},
}

var BoolDecodecs = map[string]BoolDecodec{
	"still": {Default: false},

	"hidecpu":  {Default: false},
	"hidedf":   {Default: false},
	"hideif":   {Default: false},
	"hidemem":  {Default: false},
	"hideps":   {Default: false},
	"hideswap": {Default: false},
	"hidevg":   {Default: false},

	"showconfigmem": {Default: false},
}

func (bp *BoolParam) Decode(form url.Values) {
	values, ok := form[bp.BoolDecodec.Pname]
	if !ok {
		bp.BoolDecoded.Value = bp.BoolDecodec.Default
		return
	}
	bp.BoolDecoded.Specified = true
	if len(values) != 0 || values[0] == "" || values[0] == "1" || values[0] == "true" || values[0] == "TRUE" {
		bp.BoolDecoded.Value = true
	} // else .Value stays false
	bp.Params.Values.Set(bp.BoolDecodec.Pname, bp.StringValue(bp.BoolDecoded.Value))
}

func (bp BoolParam) StringValue(value bool) string {
	if value != bp.BoolDecodec.Default {
		return ""
	}
	return fmt.Sprintf("%t", value)
}

func (ep *EnumParam) Decode(form url.Values, setep *EnumParam) error {
	_, uptr := ep.EnumDecodec.Unew()
	n, spec, err := ep.Find(form[ep.EnumDecodec.Pname], uptr)
	if err != nil {
		return err
	}
	ep.EnumDecoded.Number = n
	ep.EnumDecoded.Specified = spec
	if setep != nil {
		*setep = *ep
	}
	return nil
}

// Find side effects: uptr.Unmarshal and p.Values.Set
func (ep *EnumParam) Find(values []string, uptr Upointer) (Number, bool, error) {
	if len(values) == 0 || values[0] == "" {
		return ep.EnumDecodec.Default, false, nil
	}
	var negate bool
	in := values[0]
	if in[0] == '-' {
		in = in[1:]
		negate = true
	}
	err := uptr.Unmarshal(in, &negate) // .UnmarshalJSON([]byte(fmt.Sprintf("%q", strings.ToUpper(in))))
	if err != nil {
		if _, ok := err.(enums.RenamedConstError); ok {
			// The case when err (of type RenamedConstError) is set
			// AND uptr actually holds corresponding ("renamed") value.
			if _, l, err := uptr.Marshal(); err == nil {
				if negate {
					l = "-" + l
				}
				ep.Params.Values.Set(ep.EnumDecodec.Pname, l)
			}
			ep.Params.Moved = true
		}
		return Number{}, true, err
	}
	n := Number{
		Uint:     uptr.Touint(),
		Negative: negate,
	}
	ep.Params.Values.Set(ep.EnumDecodec.Pname, values[0])
	return n, true, nil
}

// NewParams constructs new Params.
// Global var EnumDecodecs, BoolDecodecs are ranged.
func NewParams() *Params {
	p := &Params{Values: make(url.Values)}
	enums := make(Enums)
	for k, v := range EnumDecodecs {
		v.Pname = k
		enums[k] = &EnumParam{
			EnumDecodec: v,
			Params:      p,
		}
	}
	bools := make(Bools)
	for k, v := range BoolDecodecs {
		v.Pname = k
		bools[k] = &BoolParam{
			BoolDecodec: v,
			Params:      p,
		}
	}
	p.ENUM = enums
	p.BOOL = bools
	return p
}

// EnumDecoded has Decode result.
type EnumDecoded struct {
	Number
	Specified bool
}

// EnumParam holds everything known about enum param.
// All fields are non-marshaled, there's .MarshalJSON method for that.
type EnumParam struct {
	// EnumDecodec is read-only, an entry from global var EnumDecodecs.
	EnumDecodec EnumDecodec `json:"-"` // non-marshaled explicitly
	EnumDecoded EnumDecoded `json:"-"` // non-marshaled explicitly
	Params      *Params     `json:"-"` // non-marshaled explicitly
}

func (ep EnumParam) LessorMore(r bool) bool {
	// numeric values: flip r
	if !ep.IsAlpha(ep.EnumDecoded.Number.Uint) {
		r = !r
	}
	if ep.EnumDecoded.Number.Negative {
		r = !r
	}
	return r
}

// MarshalJSON goes over all defined constants
// (by the means of p.EnumDecodec.Unew() & .Marshal method of Uinter)
// to returns a map of constants to DropLink.
func (ep EnumParam) MarshalJSON() ([]byte, error) {
	m := map[string]DropLink{}
	name, uptr := ep.EnumDecodec.Unew()
	uter := uptr.(enums.Uinter) // downcast. Upointer inlines Uinter.
	marshal := uptr.Marshal
	for i := 0; i < 100; i++ {
		nextuter, s, err := marshal()
		if err != nil {
			break
		}
		m[strings.ToUpper(s)] = ep.EncodeUint(name, uter)
		marshal = nextuter.Marshal
		uter = nextuter
	}
	return json.Marshal(m)
}

func (bp BoolParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Href  template.HTMLAttr
		Value bool
	}{
		Href:  bp.EncodeToggle(),
		Value: bp.BoolDecoded.Value,
	})
}

type Enums map[string]*EnumParam
type Bools map[string]*BoolParam

type Params struct {
	ENUM   Enums
	BOOL   Bools
	Values url.Values `json:"-"`
	Moved  bool       `json:"-"`
}

func (ps Params) ValuesCopy() url.Values {
	copy := url.Values{}
	for k, v := range ps.Values {
		copy[k] = v
	}
	return copy
}

type BoolDecodec struct {
	Default bool
	Pname   string
}

type BoolDecoded struct {
	Value     bool
	Specified bool
}

type BoolParam struct {
	// BoolDecoder is read-only, an entry from global var BoolDecoders.
	BoolDecodec BoolDecodec `json:"-"` // non marshalled explicitly
	BoolDecoded BoolDecoded `json:"-"` // non marshalled explicitly
	Params      *Params     `json:"-"` // non marshalled explicitly
}

// Href is to be removed. stays here for better diff.
type Href struct {
	Href string
}

// EncodeToggle returns template.HTMLAttr having the bp value inverted and encoded.
// The other values are copied from bp.Params.Values.
func (bp BoolParam) EncodeToggle() template.HTMLAttr {
	base := bp.Params.ValuesCopy()
	value := !bp.BoolDecoded.Value // here's the toggle
	if value == bp.BoolDecodec.Default {
		base.Del(bp.BoolDecodec.Pname)
	} else {
		base.Set(bp.BoolDecodec.Pname, bp.StringValue(value))
	}
	return template.HTMLAttr("?" + ValuesEncode(base))
}

func ValuesEncode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := url.QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			if v == "" {
				continue
			}
			buf.WriteString("=" + url.QueryEscape(v))
		}
	}
	return buf.String()
}
