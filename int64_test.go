package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64_UnmarshalJSON(t *testing.T) {
	int := new(Int64)
	err := int.UnmarshalJSON([]byte(`100`))
	assert.NoError(t, err)
	assert.EqualValues(t, 100, int.Int64)

	int = new(Int64)
	err = int.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestInt64_MarshalJSON(t *testing.T) {
	int := Int64{
		NullInt64: sql.NullInt64{Int64: 0, Valid: true},
	}
	raw, err := int.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `0`, raw)

	int.Int64 = 100
	raw, err = int.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `100`, raw)
}
