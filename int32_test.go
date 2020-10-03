package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt32_UnmarshalJSON(t *testing.T) {
	int := new(Int32)
	err := int.UnmarshalJSON([]byte(`100`))
	assert.NoError(t, err)
	assert.EqualValues(t, 100, int.Int32)

	int = new(Int32)
	err = int.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestInt32_MarshalJSON(t *testing.T) {
	int := Int32{
		NullInt32: sql.NullInt32{Int32: 0, Valid: true},
	}
	raw, err := int.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `0`, raw)

	int.Int32 = 100
	raw, err = int.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `100`, raw)
}
