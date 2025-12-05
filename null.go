package jsontypes

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Null represents an optional JSON field of type null.
type Null struct {
	// Present is true if the JSON field exists.
	Present bool
}

// IsDefined reports whether the field was present in the input JSON.
//
// It is used by easyjson to determine whether the field should be marshaled
// when using the `omitempty` tag.
func (v Null) IsDefined() bool {
	return v.Present
}

// MarshalEasyJSON implements easyjson.Marshaler.
func (v Null) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawString("null")
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler.
func (v *Null) UnmarshalEasyJSON(l *jlexer.Lexer) {
	l.Null()
	v.Present = true
}
