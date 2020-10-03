package null

import (
	"database/sql"
	"encoding/json"
)

type Time struct {
	sql.NullTime
}

func (t *Time) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &t.Time); err != nil {
		return err
	}
	t.Valid = true
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}
