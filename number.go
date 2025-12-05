package jsontypes

import (
	"encoding/json"
	"strconv"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type numberKind byte

const (
	kindString = numberKind(iota)
	kindInt
	kindInt8
	kindInt16
	kindInt32
	kindInt64
	kindUInt
	kindUInt8
	kindUInt16
	kindUInt32
	kindUInt64
	kindFloat32
	kindFloat64
)

// Number represents an optional JSON field of type number.
//
// It models three distinct states:
//   - field not present in the JSON:     Present = false, Valid = false
//   - field present with null value:     Present = true,  Valid = false
//   - field present with a real value:   Present = true,  Valid = true
//
// This is useful when you need to know whether a field existed in the input,
// not just whether its value is null.
type Number struct {
	kind numberKind

	// Present is true if the JSON field exists, even if the value is null.
	Present bool

	// Valid is true only when the field is present and the value is non-null.
	Valid bool

	// Value holds the underlying value when both Present and Valid are true.
	Value json.Number

	signed   int64
	unsigned uint64
	float    float64
}

// IsDefined reports whether the field was present in the input JSON,
// regardless of whether it contained null or a non-null value.
//
// It is used by easyjson to determine whether the field should be marshaled
// when using the `omitempty` tag.
func (v *Number) IsDefined() bool {
	return v.Present
}

// Get returns the contained value if the field is present and non-null.
// Otherwise, it returns the supplied fallback value.
func (v *Number) Get(value json.Number) json.Number {
	if v.Present && v.Valid && len(v.Value) > 0 {
		return v.Value
	} else {
		return value
	}
}

// Set assigns a non-null value and marks the field as present.
func (v *Number) Set(value json.Number) {
	v.Present = true
	v.Valid = true
	v.Value = value
	v.kind = kindString
}

// MarshalEasyJSON implements easyjson.Marshaler.
func (v *Number) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Valid {
		switch v.kind {
		case kindString:
			w.RawString(string(v.Value))
		case kindInt:
			w.Int(int(v.signed))
		case kindInt8:
			w.Int8(int8(v.signed))
		case kindInt16:
			w.Int16(int16(v.signed))
		case kindInt32:
			w.Int32(int32(v.signed))
		case kindInt64:
			w.Int64(v.signed)
		case kindUInt:
			w.Uint(uint(v.unsigned))
		case kindUInt8:
			w.Uint8(uint8(v.unsigned))
		case kindUInt16:
			w.Uint16(uint16(v.unsigned))
		case kindUInt32:
			w.Uint32(uint32(v.unsigned))
		case kindUInt64:
			w.Uint64(v.unsigned)
		case kindFloat32:
			w.Float32(float32(v.float))
		case kindFloat64:
			w.Float64(v.float)
		default:
			panic("cannot marshal unknown number kind")
		}
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler.
func (v *Number) UnmarshalEasyJSON(l *jlexer.Lexer) {
	v.Present = true
	if l.IsNull() {
		l.Skip()
	} else {
		v.Valid = true
		v.Value = l.JsonNumber()
	}
}

// ParseInt parses the underlying number value to an int.
func (v *Number) ParseInt() (int, error) {
	if v.kind != kindInt {
		var err error
		if v.signed, err = strconv.ParseInt(string(v.Value), 10, strconv.IntSize); err == nil {
			v.kind = kindInt
		} else {
			return 0, err
		}
	}
	return int(v.signed), nil
}

// Int returns the number as an int.
// Returns zero in case of parse error.
func (v *Number) Int() int {
	if v.kind != kindInt {
		value, _ := v.ParseInt()
		return value
	}
	return int(v.signed)
}

// ParseInt8 parses the underlying number value to an int8.
func (v *Number) ParseInt8() (int8, error) {
	if v.kind != kindInt8 {
		var err error
		if v.signed, err = strconv.ParseInt(string(v.Value), 10, 8); err == nil {
			v.kind = kindInt8
		} else {
			return 0, err
		}
	}
	return int8(v.signed), nil
}

// Int8 returns the number as an int8.
// Returns zero in case of parse error.
func (v *Number) Int8() int8 {
	if v.kind != kindInt8 {
		value, _ := v.ParseInt8()
		return value
	}
	return int8(v.signed)
}

// ParseInt16 parses the underlying number value to an int16.
func (v *Number) ParseInt16() (int16, error) {
	if v.kind != kindInt16 {
		var err error
		if v.signed, err = strconv.ParseInt(string(v.Value), 10, 16); err == nil {
			v.kind = kindInt16
		} else {
			return 0, err
		}
	}
	return int16(v.signed), nil
}

// Int16 returns the number as an int16.
// Returns zero in case of parse error.
func (v *Number) Int16() int16 {
	if v.kind != kindInt16 {
		value, _ := v.ParseInt16()
		return value
	}
	return int16(v.signed)
}

// ParseInt32 parses the underlying number value to an int32.
func (v *Number) ParseInt32() (int32, error) {
	if v.kind != kindInt32 {
		var err error
		if v.signed, err = strconv.ParseInt(string(v.Value), 10, 32); err == nil {
			v.kind = kindInt32
		} else {
			return 0, err
		}
	}
	return int32(v.signed), nil
}

// Int32 returns the number as an int32.
// Returns zero in case of parse error.
func (v *Number) Int32() int32 {
	if v.kind != kindInt32 {
		value, _ := v.ParseInt32()
		return value
	}
	return int32(v.signed)
}

// ParseInt64 parses the underlying number value to an int64.
func (v *Number) ParseInt64() (int64, error) {
	if v.kind != kindInt64 {
		var err error
		if v.signed, err = strconv.ParseInt(string(v.Value), 10, 64); err == nil {
			v.kind = kindInt64
		} else {
			return 0, err
		}
	}
	return v.signed, nil
}

// Int64 returns the number as an int64.
// Returns zero in case of parse error.
func (v *Number) Int64() int64 {
	if v.kind != kindInt64 {
		value, _ := v.ParseInt64()
		return value
	}
	return v.signed
}

// ParseUInt parses the underlying number value to an uint.
func (v *Number) ParseUInt() (uint, error) {
	if v.kind != kindUInt {
		var err error
		if v.unsigned, err = strconv.ParseUint(string(v.Value), 10, strconv.IntSize); err == nil {
			v.kind = kindUInt
		} else {
			return 0, err
		}
	}
	return uint(v.unsigned), nil
}

// UInt returns the number as an uint.
// Returns zero in case of parse error.
func (v *Number) UInt() uint {
	if v.kind != kindUInt {
		value, _ := v.ParseUInt()
		return value
	}
	return uint(v.unsigned)
}

// ParseUInt8 parses the underlying number value to an uint8.
func (v *Number) ParseUInt8() (uint8, error) {
	if v.kind != kindUInt8 {
		var err error
		if v.unsigned, err = strconv.ParseUint(string(v.Value), 10, 8); err == nil {
			v.kind = kindUInt8
		} else {
			return 0, err
		}
	}
	return uint8(v.unsigned), nil
}

// UInt8 returns the number as an uint8.
// Returns zero in case of parse error.
func (v *Number) UInt8() uint8 {
	if v.kind != kindUInt8 {
		value, _ := v.ParseUInt8()
		return value
	}
	return uint8(v.unsigned)
}

// ParseUInt16 parses the underlying number value to an uint16.
func (v *Number) ParseUInt16() (uint16, error) {
	if v.kind != kindUInt16 {
		var err error
		if v.unsigned, err = strconv.ParseUint(string(v.Value), 10, 16); err == nil {
			v.kind = kindUInt16
		} else {
			return 0, err
		}
	}
	return uint16(v.unsigned), nil
}

// UInt16 returns the number as an uint16.
// Returns zero in case of parse error.
func (v *Number) UInt16() uint16 {
	if v.kind != kindUInt16 {
		value, _ := v.ParseUInt16()
		return value
	}
	return uint16(v.unsigned)
}

// ParseUInt32 parses the underlying number value to an uint32.
func (v *Number) ParseUInt32() (uint32, error) {
	if v.kind != kindUInt32 {
		var err error
		if v.unsigned, err = strconv.ParseUint(string(v.Value), 10, 32); err == nil {
			v.kind = kindUInt32
		} else {
			return 0, err
		}
	}
	return uint32(v.unsigned), nil
}

// UInt32 returns the number as an uint32.
// Returns zero in case of parse error.
func (v *Number) UInt32() uint32 {
	if v.kind != kindUInt32 {
		value, _ := v.ParseUInt32()
		return value
	}
	return uint32(v.unsigned)
}

// ParseUInt64 parses the underlying number value to an uint64.
func (v *Number) ParseUInt64() (uint64, error) {
	if v.kind != kindUInt64 {
		var err error
		if v.unsigned, err = strconv.ParseUint(string(v.Value), 10, 64); err == nil {
			v.kind = kindUInt64
		} else {
			return 0, err
		}
	}
	return v.unsigned, nil
}

// UInt64 returns the number as an uint64.
// Returns zero in case of parse error.
func (v *Number) UInt64() uint64 {
	if v.kind != kindUInt64 {
		value, _ := v.ParseUInt64()
		return value
	}
	return v.unsigned
}

// ParseFloat32 parses the underlying number value to a float32.
func (v *Number) ParseFloat32() (float32, error) {
	if v.kind != kindFloat32 {
		var err error
		if v.float, err = strconv.ParseFloat(string(v.Value), 32); err == nil {
			v.kind = kindFloat32
		} else {
			return 0, err
		}
	}
	return float32(v.float), nil
}

// Float32 returns the number as a float32.
// Returns zero in case of parse error.
func (v *Number) Float32() float32 {
	if v.kind != kindFloat32 {
		value, _ := v.ParseFloat32()
		return value
	}
	return float32(v.float)
}

// ParseFloat64 parses the underlying number value to a float64.
func (v *Number) ParseFloat64() (float64, error) {
	if v.kind != kindFloat64 {
		var err error
		if v.float, err = strconv.ParseFloat(string(v.Value), 64); err == nil {
			v.kind = kindFloat64
		} else {
			return 0, err
		}
	}
	return v.float, nil
}

// Float64 returns the number as a float64.
// Returns zero in case of parse error.
func (v *Number) Float64() float64 {
	if v.kind != kindFloat64 {
		value, _ := v.ParseFloat64()
		return value
	}
	return v.float
}

// SetInt assigns an int as the underlying number value.
func (v *Number) SetInt(value int) {
	v.kind = kindInt
	v.signed = int64(value)
}

// SetInt8 assigns an int8 as the underlying number value.
func (v *Number) SetInt8(value int8) {
	v.kind = kindInt8
	v.signed = int64(value)
}

// SetInt16 assigns an int16 as the underlying number value.
func (v *Number) SetInt16(value int16) {
	v.kind = kindInt16
	v.signed = int64(value)
}

// SetInt32 assigns an int32 as the underlying number value.
func (v *Number) SetInt32(value int32) {
	v.kind = kindInt32
	v.signed = int64(value)
}

// SetInt64 assigns an int64 as the underlying number value.
func (v *Number) SetInt64(value int64) {
	v.kind = kindInt64
	v.signed = value
}

// SetUInt assigns an uint as the underlying number value.
func (v *Number) SetUInt(value uint) {
	v.kind = kindUInt
	v.unsigned = uint64(value)
}

// SetUInt8 assigns an uint8 as the underlying number value.
func (v *Number) SetUInt8(value uint8) {
	v.kind = kindUInt8
	v.unsigned = uint64(value)
}

// SetUInt16 assigns an uint16 as the underlying number value.
func (v *Number) SetUInt16(value uint16) {
	v.kind = kindUInt16
	v.unsigned = uint64(value)
}

// SetUInt32 assigns an uint32 as the underlying number value.
func (v *Number) SetUInt32(value uint32) {
	v.kind = kindUInt32
	v.unsigned = uint64(value)
}

// SetUInt64 assigns an uint64 as the underlying number value.
func (v *Number) SetUInt64(value uint64) {
	v.kind = kindUInt64
	v.unsigned = value
}

// SetFloat32 assigns a float32 as the underlying number value.
func (v *Number) SetFloat32(value float32) {
	v.kind = kindFloat32
	v.float = float64(value)
}

// SetFloat64 assigns a float64 as the underlying number value.
func (v *Number) SetFloat64(value float64) {
	v.kind = kindFloat64
	v.float = value
}
