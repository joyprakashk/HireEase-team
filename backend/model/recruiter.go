package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Recruiter struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Company   string      `json:"company"`
	Jobs      []uuid.UUID `json:"jobs"`
	Password  string      `json:"password,omitempty"`
	CreatedAt string      `json:"-"`
	UpdatedAt string      `json:"-"`
	DeletedAt string      `json:"-"`
}

func (r Recruiter) Value() (driver.Value, error) {
	r.Password = ""
	return json.Marshal(r)
}

func (r *Recruiter) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &r)
}
