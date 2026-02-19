// Package apijson provides Field[T], a generic optional/nullable field type
// for use in API request and response structs.
package apijson

import "encoding/json"

// Field[T] represents a JSON field that can be in one of three states:
//   - not present (omitted from JSON output)
//   - explicitly null ("null" in JSON)
//   - present with a value
type Field[T any] struct {
	Value   T
	null    bool
	present bool
}

// F creates a Field[T] with the given value (present, not null).
func F[T any](v T) Field[T] {
	return Field[T]{Value: v, present: true}
}

// Null creates a Field[T] that serializes as JSON null.
func Null[T any]() Field[T] {
	return Field[T]{null: true, present: true}
}

// IsPresent reports whether the field was set (either to a value or to null).
func (f Field[T]) IsPresent() bool { return f.present }

// IsNull reports whether the field is explicitly null.
func (f Field[T]) IsNull() bool { return f.null }

// MarshalJSON implements json.Marshaler.
// - If not present and not null: returns nil (field is omitted when tagged omitempty).
// - If null: returns "null".
// - Otherwise: marshals Value.
func (f Field[T]) MarshalJSON() ([]byte, error) {
	if !f.present && !f.null {
		return []byte("null"), nil
	}
	if f.null {
		return []byte("null"), nil
	}
	return json.Marshal(f.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *Field[T]) UnmarshalJSON(data []byte) error {
	f.present = true
	if string(data) == "null" {
		f.null = true
		return nil
	}
	f.null = false
	return json.Unmarshal(data, &f.Value)
}
