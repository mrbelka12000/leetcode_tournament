package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Points struct {
	Easy   uint64 `json:"easy,omitempty"`
	Medium uint64 `json:"medium,omitempty"`
	Hard   uint64 `json:"hard,omitempty"`
	Total  uint64 `json:"total,omitempty"`
}

// Scan method to unmarshal jsonb from postgres
func (p *Points) Scan(src interface{}) (err error) {
	switch src.(type) {
	case []byte:
		err = json.Unmarshal(src.([]byte), &p)
	default:
		return errors.New("Incompatible type for Skills")
	}
	if err != nil {
		return
	}
	return nil
}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (p Points) Value() (driver.Value, error) {
	return json.Marshal(p)
}
