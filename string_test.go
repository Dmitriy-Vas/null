package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_UnmarshalJSON(t *testing.T) {
	string := new(String)
	err := string.UnmarshalJSON([]byte(`"hello world"`))
	assert.NoError(t, err)
	assert.EqualValues(t, `hello world`, string.String)

	string = new(String)
	err = string.UnmarshalJSON([]byte(`""`))
	assert.NoError(t, err)
	assert.EqualValues(t, ``, string.String)

	string = new(String)
	err = string.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestString_MarshalJSON(t *testing.T) {
	string := String{
		NullString: sql.NullString{String: `%Hello World%`, Valid: true},
	}
	raw, err := string.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `"%Hello World%"`, raw)

	string.String = `Hello World`
	raw, err = string.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `"Hello World"`, raw)
}
