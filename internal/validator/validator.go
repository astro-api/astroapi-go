// Package validator provides simple struct validation via reflection
// and struct tags.
package validator

import (
	"fmt"
	"reflect"
	"strings"
)

// Validatable is an optional interface that types can implement to
// provide custom validation logic in addition to tag-based validation.
type Validatable interface {
	Validate() error
}

// Validate checks params for fields tagged with `validate:"required"`.
// It also calls params.Validate() if the Validatable interface is implemented.
// params may be nil or a pointer to nil — both are safe.
func Validate(params any) error {
	if params == nil {
		return nil
	}

	v := reflect.ValueOf(params)
	// Dereference pointer(s).
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	if err := checkRequired(v); err != nil {
		return err
	}

	if val, ok := params.(Validatable); ok {
		return val.Validate()
	}

	return nil
}

func checkRequired(v reflect.Value) error {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

		tag := field.Tag.Get("validate")
		if !hasRequired(tag) {
			continue
		}

		if isEmpty(val) {
			name := fieldName(field)
			return fmt.Errorf("field %q is required", name)
		}
	}
	return nil
}

func hasRequired(tag string) bool {
	for _, part := range strings.Split(tag, ",") {
		if strings.TrimSpace(part) == "required" {
			return true
		}
	}
	return false
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return false // booleans can't really be "missing"
	case reflect.Struct:
		return v.IsZero()
	default:
		return false
	}
}

func fieldName(f reflect.StructField) string {
	if tag := f.Tag.Get("json"); tag != "" {
		name := strings.Split(tag, ",")[0]
		if name != "" && name != "-" {
			return name
		}
	}
	return f.Name
}
