package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	value := new(Time)
	err := value.UnmarshalJSON([]byte(`"1970-01-01T05:00:00+05:00"`))
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		time.Unix(0, 0),
		value.Time)

	value = new(Time)
	err = value.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestTime_MarshalJSON(t *testing.T) {
	value := Time{
		NullTime: sql.NullTime{
			Time:  time.Unix(0, 0),
			Valid: true},
	}
	raw, err := value.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `"1970-01-01T05:00:00+05:00"`, string(raw))
}
