package null

import (
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
}

func (s *String) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &s.String); err != nil {
		return err
	}
	s.Valid = true
	return nil
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String)
}
