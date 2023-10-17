package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Points struct {
	Easy   uint64 `json:"easy" schema:"easy"`
	Medium uint64 `json:"medium" schema:"medium"`
	Hard   uint64 `json:"hard" schema:"hard"`
	Total  uint64 `json:"total" schema:"total"`
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
func (p Points) Value() (driver.Value, error) {
	return json.Marshal(p)
}
