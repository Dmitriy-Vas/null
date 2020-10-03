package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64_UnmarshalJSON(t *testing.T) {
	value := new(Int64)
	err := value.UnmarshalJSON([]byte(`100`))
	assert.NoError(t, err)
	assert.EqualValues(t, 100, value.Int64)

	value = new(Int64)
	err = value.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestInt64_MarshalJSON(t *testing.T) {
	value := Int64{
		NullInt64: sql.NullInt64{Int64: 0, Valid: true},
	}
	raw, err := value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `0`, raw)

	value.Int64 = 100
	raw, err = value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `100`, raw)
}
