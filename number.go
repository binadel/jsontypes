package jsontypes

import (
	"encoding/json"
	"strconv"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type numberKind byte

const (
	kindInt = numberKind(iota)
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
	// Present is true if the JSON field exists, even if the value is null.
	Present bool

	// Valid is true only when the field is present and the value is non-null.
	Valid bool

	// Value holds the underlying value when both Present and Valid are true.
	Value json.Number

	kind    numberKind
	integer int64
	float   float64
}

// IsDefined reports whether the field was present in the input JSON,
// regardless of whether it contained null or a non-null value.
//
// It is used by easyjson to determine whether the field should be marshaled
// when using the `omitempty` tag.
func (v Number) IsDefined() bool {
	return v.Present
}

// Get returns the contained value if the field is present and non-null.
// Otherwise, it returns the supplied fallback value.
func (v Number) Get(value json.Number) json.Number {
	if v.Present && v.Valid {
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
}

// MarshalEasyJSON implements easyjson.Marshaler.
func (v Number) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Valid {
		w.RawString(string(v.Value))
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

// Int parses the underlying number value to an int type.
func (v Number) Int() (int, error) {
	n, err := strconv.ParseInt(string(v.Value), 10, strconv.IntSize)
	return int(n), err
}

// Int8 parses the underlying number value to an int8 type.
func (v Number) Int8() (int8, error) {
	n, err := strconv.ParseInt(string(v.Value), 10, 8)
	return int8(n), err
}

// Int16 parses the underlying number value to an int16 type.
func (v Number) Int16() (int16, error) {
	n, err := strconv.ParseInt(string(v.Value), 10, 16)
	return int16(n), err
}

// Int32 parses the underlying number value to an int32 type.
func (v Number) Int32() (int32, error) {
	n, err := strconv.ParseInt(string(v.Value), 10, 32)
	return int32(n), err
}

// Int64 parses the underlying number value to an int64 type.
func (v Number) Int64() (int64, error) {
	return strconv.ParseInt(string(v.Value), 10, 64)
}

// UInt parses the underlying number value to an uint type.
func (v Number) UInt() (uint, error) {
	n, err := strconv.ParseUint(string(v.Value), 10, strconv.IntSize)
	return uint(n), err
}

// UInt8 parses the underlying number value to an uint8 type.
func (v Number) UInt8() (uint8, error) {
	n, err := strconv.ParseUint(string(v.Value), 10, 8)
	return uint8(n), err
}

// UInt16 parses the underlying number value to an uint16 type.
func (v Number) UInt16() (uint16, error) {
	n, err := strconv.ParseUint(string(v.Value), 10, 16)
	return uint16(n), err
}

// UInt32 parses the underlying number value to an uint32 type.
func (v Number) UInt32() (uint32, error) {
	n, err := strconv.ParseUint(string(v.Value), 10, 32)
	return uint32(n), err
}

// UInt64 parses the underlying number value to an uint64 type.
func (v Number) UInt64() (uint64, error) {
	return strconv.ParseUint(string(v.Value), 10, 64)
}

// Float32 parses the underlying number value to a float32 type.
func (v Number) Float32() (float32, error) {
	n, err := strconv.ParseFloat(string(v.Value), 32)
	return float32(n), err
}

// Float64 parses the underlying number value to a float64 type.
func (v Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(v.Value), 64)
}

// SetInt formats the number value from an int type.
func (v *Number) SetInt(value int) {
	v.kind = kindInt
	v.integer = int64(value)
}

// SetInt8 formats the number value from an int8 type.
func (v *Number) SetInt8(value int8) {
	v.kind = kindInt8
	v.integer = int64(value)
}

// SetInt16 formats the number value from an int16 type.
func (v *Number) SetInt16(value int16) {
	v.kind = kindInt16
	v.integer = int64(value)
}

// SetInt32 formats the number value from an int32 type.
func (v *Number) SetInt32(value int32) {
	v.kind = kindInt32
	v.integer = int64(value)
}

// SetInt64 formats the number value from an int64 type.
func (v *Number) SetInt64(value int64) {
	v.kind = kindInt64
	v.integer = value
}

// SetUInt formats the number value from an uint type.
func (v *Number) SetUInt(value uint) {
	v.kind = kindUInt
	v.integer = int64(value)
}

// SetUInt8 formats the number value from an uint8 type.
func (v *Number) SetUInt8(value uint8) {
	v.kind = kindUInt8
	v.integer = int64(value)
}

// SetUInt16 formats the number value from an uint16 type.
func (v *Number) SetUInt16(value uint16) {
	v.kind = kindUInt16
	v.integer = int64(value)
}

// SetUInt32 formats the number value from an uint32 type.
func (v *Number) SetUInt32(value uint32) {
	v.kind = kindUInt32
	v.integer = int64(value)
}

// SetUInt64 formats the number value from an uint64 type.
func (v *Number) SetUInt64(value uint64) {
	v.kind = kindUInt64
	v.integer = int64(value)
}

// SetFloat32 formats the number value from a float32 type.
func (v *Number) SetFloat32(value float32) {
	v.kind = kindFloat32
	v.float = float64(value)
}

// SetFloat64 formats the number value from a float64 type.
func (v *Number) SetFloat64(value float64) {
	v.kind = kindFloat64
	v.float = value
}
