package null

import (
	"database/sql"
	"encoding/json"
)

type Float64 struct {
	sql.NullFloat64
}

func (f *Float64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &f.Float64); err != nil {
		return err
	}
	f.Valid = true
	return nil
}

func (f *Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Float64)
}
