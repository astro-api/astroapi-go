package validator_test

import (
	"fmt"
	"testing"

	"github.com/astro-api/astroapi-go/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type requiredStringParams struct {
	Name string `json:"name" validate:"required"`
	Note string `json:"note"`
}

type requiredIntParams struct {
	Year int `json:"year" validate:"required"`
}

type requiredPtrParams struct {
	Subject *struct{ Name string } `json:"subject" validate:"required"`
}

func TestValidate_NilIsOK(t *testing.T) {
	err := validator.Validate(nil)
	assert.NoError(t, err)
}

func TestValidate_MissingRequiredString(t *testing.T) {
	p := &requiredStringParams{}
	err := validator.Validate(p)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "name")
}

func TestValidate_PresentRequiredString(t *testing.T) {
	p := &requiredStringParams{Name: "Alice"}
	err := validator.Validate(p)
	assert.NoError(t, err)
}

func TestValidate_OptionalFieldCanBeEmpty(t *testing.T) {
	p := &requiredStringParams{Name: "Alice"}
	err := validator.Validate(p)
	assert.NoError(t, err)
}

func TestValidate_MissingRequiredInt(t *testing.T) {
	p := &requiredIntParams{}
	err := validator.Validate(p)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "year")
}

func TestValidate_PresentRequiredInt(t *testing.T) {
	p := &requiredIntParams{Year: 1990}
	err := validator.Validate(p)
	assert.NoError(t, err)
}

func TestValidate_MissingRequiredPointer(t *testing.T) {
	p := &requiredPtrParams{}
	err := validator.Validate(p)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "subject")
}

type customValidatable struct {
	Name string
}

func (c customValidatable) Validate() error {
	if c.Name == "bad" {
		return fmt.Errorf("name cannot be 'bad'")
	}
	return nil
}

func TestValidate_CustomValidatable(t *testing.T) {
	p := customValidatable{Name: "bad"}
	err := validator.Validate(p)
	require.Error(t, err)
}

func TestValidate_CustomValidatable_OK(t *testing.T) {
	p := customValidatable{Name: "good"}
	err := validator.Validate(p)
	assert.NoError(t, err)
}

func TestValidate_NonStruct(t *testing.T) {
	err := validator.Validate("just a string")
	assert.NoError(t, err)
}
