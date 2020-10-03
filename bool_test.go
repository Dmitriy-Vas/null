package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	bool := new(Bool)
	err := bool.UnmarshalJSON([]byte(`false`))
	assert.NoError(t, err)
	assert.EqualValues(t, false, bool.Bool)

	bool = new(Bool)
	err = bool.UnmarshalJSON([]byte(`true`))
	assert.NoError(t, err)
	assert.EqualValues(t, true, bool.Bool)

	bool = new(Bool)
	err = bool.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestBool_MarshalJSON(t *testing.T) {
	bool := Bool{
		NullBool: sql.NullBool{Bool: true, Valid: true},
	}
	raw, err := bool.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `true`, raw)

	bool.Bool = false
	raw, err = bool.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `false`, raw)

	bool.Valid = false
	raw, err = bool.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `false`, raw)
}
