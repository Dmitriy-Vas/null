package null

import (
	"database/sql"
	"encoding/json"
)

type Int32 struct {
	sql.NullInt32
}

func (i *Int32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.Int32); err != nil {
		return err
	}
	i.Valid = true
	return nil
}

func (i *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Int32)
}
