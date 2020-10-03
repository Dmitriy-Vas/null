package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	value := new(Bool)
	err := value.UnmarshalJSON([]byte(`false`))
	assert.NoError(t, err)
	assert.EqualValues(t, false, value.Bool)

	value = new(Bool)
	err = value.UnmarshalJSON([]byte(`true`))
	assert.NoError(t, err)
	assert.EqualValues(t, true, value.Bool)

	value = new(Bool)
	err = value.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestBool_MarshalJSON(t *testing.T) {
	value := Bool{
		NullBool: sql.NullBool{Bool: true, Valid: true},
	}
	raw, err := value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `true`, raw)

	value.Bool = false
	raw, err = value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `false`, raw)

	value.Valid = false
	raw, err = value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `false`, raw)
}
