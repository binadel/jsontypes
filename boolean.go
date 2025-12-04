package jsontypes

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Boolean represents an optional JSON field of type boolean.
//
// It models three distinct states:
//   - field not present in the JSON:     Present = false, Valid = false
//   - field present with null value:     Present = true,  Valid = false
//   - field present with a real value:   Present = true,  Valid = true
//
// This is useful when you need to know whether a field existed in the input,
// not just whether its value is null.
type Boolean struct {
	// Present is true if the JSON field exists, even if the value is null.
	Present bool

	// Valid is true only when the field is present and the value is non-null.
	Valid bool

	// Value holds the underlying bool when both Present and Valid are true.
	Value bool
}

// IsDefined reports whether the field was present in the input JSON,
// regardless of whether it contained null or a non-null value.
//
// It is used by easyjson to determine whether the field should be marshaled
// when using the `omitempty` tag.
func (v Boolean) IsDefined() bool {
	return v.Present
}

// Get returns the contained value if the field is present and non-null.
// Otherwise, it returns the supplied fallback value.
func (v Boolean) Get(value bool) bool {
	if v.Present && v.Valid {
		return v.Value
	} else {
		return value
	}
}

// Set assigns a non-null value and marks the field as present.
func (v *Boolean) Set(value bool) {
	v.Present = true
	v.Valid = true
	v.Value = value
}

// MarshalEasyJSON implements easyjson.Marshaler.
func (v Boolean) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Valid {
		w.Bool(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler.
func (v *Boolean) UnmarshalEasyJSON(l *jlexer.Lexer) {
	v.Present = true
	if l.IsNull() {
		l.Skip()
	} else {
		v.Valid = true
		v.Value = l.Bool()
	}
}
