package null

import (
	"database/sql"
	"encoding/json"
)

type Int64 struct {
	sql.NullInt64
}

func (i *Int64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.Int64); err != nil {
		return err
	}
	i.Valid = true
	return nil
}

func (i *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Int64)
}
