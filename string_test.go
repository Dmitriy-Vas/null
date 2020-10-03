package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_UnmarshalJSON(t *testing.T) {
	value := new(String)
	err := value.UnmarshalJSON([]byte(`"hello world"`))
	assert.NoError(t, err)
	assert.EqualValues(t, `hello world`, value.String)

	value = new(String)
	err = value.UnmarshalJSON([]byte(`""`))
	assert.NoError(t, err)
	assert.EqualValues(t, ``, value.String)

	value = new(String)
	err = value.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestString_MarshalJSON(t *testing.T) {
	value := String{
		NullString: sql.NullString{String: `%Hello World%`, Valid: true},
	}
	raw, err := value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `"%Hello World%"`, raw)

	value.String = `Hello World`
	raw, err = value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `"Hello World"`, raw)
}
