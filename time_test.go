package null

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	tme := new(Time)
	err := tme.UnmarshalJSON([]byte(`"1970-01-01T05:00:00+05:00"`))
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		time.Unix(0, 0),
		tme.Time)

	tme = new(Time)
	err = tme.UnmarshalJSON([]byte(``))
	assert.Errorf(t, err, `unexpected end of JSON input`)
}

func TestTime_MarshalJSON(t *testing.T) {
	time := Time{
		NullTime: sql.NullTime{
			Time:  time.Unix(0, 0),
			Valid: true},
	}
	raw, err := time.MarshalJSON()
	assert.NoError(t, err)
	assert.EqualValues(t, `"1970-01-01T05:00:00+05:00"`, string(raw))
}
