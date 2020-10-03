package null

import (
	"database/sql"
	"encoding/json"
)

type Bool struct {
	sql.NullBool
}

func (b *Bool) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &b.Bool); err != nil {
		return err
	}
	b.Valid = true
	return nil
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid || !b.Bool {
		return []byte("false"), nil
	}
	return []byte("true"), nil
}
