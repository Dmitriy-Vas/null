package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64_UnmarshalJSON(t *testing.T) {
	float := new(Float64)
	err := float.UnmarshalJSON([]byte(`3.141`))
	assert.NoError(t, err)
	assert.EqualValues(t, 3.141, float.Float64)

	float = new(Float64)
	err = float.UnmarshalJSON([]byte(`3`))
	assert.NoError(t, err)
	assert.EqualValues(t, 3, float.Float64)

	float = new(Float64)
	err = float.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestFloat64_MarshalJSON(t *testing.T) {
	float := Float64{
		NullFloat64: sql.NullFloat64{Float64: 0.0, Valid: true},
	}
	raw, err := float.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `0`, raw)

	float.Float64 = 3.141
	raw, err = float.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `3.141`, raw)
}
