package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Job struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"description"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Domain    string    `json:"domain"`
	Skills    string    `json:"skills"`
	CreatedAt string    `json:"-"`
	UpdatedAt string    `json:"-"`
	DeletedAt string    `json:"-"`
}

func (j Job) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Job) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &j)
}
