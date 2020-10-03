package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt32_UnmarshalJSON(t *testing.T) {
	value := new(Int32)
	err := value.UnmarshalJSON([]byte(`100`))
	assert.NoError(t, err)
	assert.EqualValues(t, 100, value.Int32)

	value = new(Int32)
	err = value.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestInt32_MarshalJSON(t *testing.T) {
	value := Int32{
		NullInt32: sql.NullInt32{Int32: 0, Valid: true},
	}
	raw, err := value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `0`, raw)

	value.Int32 = 100
	raw, err = value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `100`, raw)
}
