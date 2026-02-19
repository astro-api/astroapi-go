package apijson_test

import (
	"encoding/json"
	"testing"

	"github.com/astro-api/astroapi-go/internal/apijson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestField_F(t *testing.T) {
	f := apijson.F("hello")
	assert.True(t, f.IsPresent())
	assert.False(t, f.IsNull())
	assert.Equal(t, "hello", f.Value)
}

func TestField_Null(t *testing.T) {
	f := apijson.Null[string]()
	assert.True(t, f.IsPresent())
	assert.True(t, f.IsNull())
}

func TestField_Zero(t *testing.T) {
	var f apijson.Field[string]
	assert.False(t, f.IsPresent())
	assert.False(t, f.IsNull())
}

func TestField_MarshalJSON_Value(t *testing.T) {
	f := apijson.F(42)
	data, err := json.Marshal(f)
	require.NoError(t, err)
	assert.Equal(t, "42", string(data))
}

func TestField_MarshalJSON_Null(t *testing.T) {
	f := apijson.Null[string]()
	data, err := json.Marshal(f)
	require.NoError(t, err)
	assert.Equal(t, "null", string(data))
}

func TestField_MarshalJSON_NotPresent(t *testing.T) {
	var f apijson.Field[string]
	data, err := json.Marshal(f)
	require.NoError(t, err)
	assert.Equal(t, "null", string(data))
}

func TestField_UnmarshalJSON_Value(t *testing.T) {
	var f apijson.Field[int]
	err := json.Unmarshal([]byte("123"), &f)
	require.NoError(t, err)
	assert.True(t, f.IsPresent())
	assert.False(t, f.IsNull())
	assert.Equal(t, 123, f.Value)
}

func TestField_UnmarshalJSON_Null(t *testing.T) {
	var f apijson.Field[string]
	err := json.Unmarshal([]byte("null"), &f)
	require.NoError(t, err)
	assert.True(t, f.IsPresent())
	assert.True(t, f.IsNull())
}

func TestField_InStruct(t *testing.T) {
	type payload struct {
		Name apijson.Field[string] `json:"name,omitempty"`
		Age  apijson.Field[int]    `json:"age,omitempty"`
	}

	p := payload{
		Name: apijson.F("Alice"),
	}

	data, err := json.Marshal(p)
	require.NoError(t, err)

	var out payload
	err = json.Unmarshal(data, &out)
	require.NoError(t, err)
	assert.Equal(t, "Alice", out.Name.Value)
}
