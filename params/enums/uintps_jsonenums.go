// generated by jsonenums -type=UintPS; DO NOT EDIT

package enums

import (
	"encoding/json"
	"fmt"
)

var (
	_UintPSNameToValue = map[string]UintPS{
		"PID":  PID,
		"PRI":  PRI,
		"NICE": NICE,
		"VIRT": VIRT,
		"RES":  RES,
		"TIME": TIME,
		"NAME": NAME,
		"UID":  UID,
		"USER": USER,
	}

	_UintPSValueToName = map[UintPS]string{
		PID:  "PID",
		PRI:  "PRI",
		NICE: "NICE",
		VIRT: "VIRT",
		RES:  "RES",
		TIME: "TIME",
		NAME: "NAME",
		UID:  "UID",
		USER: "USER",
	}
)

func init() {
	var v UintPS
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_UintPSNameToValue = map[string]UintPS{
			interface{}(PID).(fmt.Stringer).String():  PID,
			interface{}(PRI).(fmt.Stringer).String():  PRI,
			interface{}(NICE).(fmt.Stringer).String(): NICE,
			interface{}(VIRT).(fmt.Stringer).String(): VIRT,
			interface{}(RES).(fmt.Stringer).String():  RES,
			interface{}(TIME).(fmt.Stringer).String(): TIME,
			interface{}(NAME).(fmt.Stringer).String(): NAME,
			interface{}(UID).(fmt.Stringer).String():  UID,
			interface{}(USER).(fmt.Stringer).String(): USER,
		}
	}
}

// MarshalJSON is generated so UintPS satisfies json.Marshaler.
func (r UintPS) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _UintPSValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid UintPS: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so UintPS satisfies json.Unmarshaler.
func (r *UintPS) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("UintPS should be a string, got %s", data)
	}
	v, ok := _UintPSNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid UintPS %q", s)
	}
	*r = v
	return nil
}
