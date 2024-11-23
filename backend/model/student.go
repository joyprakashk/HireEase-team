package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Branch    string `json:"branch"`
	CGPA      string `json:"cgpa"`
	Courses   string `json:"courses"`
	GitHub    string `json:"github"`
	LinkedIn  string `json:"linkedin"`
	Interests string `json:"interests"`
	Skills    string `json:"skills"`
	Location  string `json:"location"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

func (s Student) Value() (driver.Value, error) {
	s.Password = ""
	return json.Marshal(s)
}

func (s *Student) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}
