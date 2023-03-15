package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

// Todo is used by pop to map your todoes database table to your go code.
type Todo struct {
	ID        int       `json:"id" db:"id"`
	Item      string    `json:"item" db:"item"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// TableName overrides the table name used by Pop.
func (u Todo) TableName() string {
	return "todos"
}

// String is not required by pop and may be deleted
func (t Todo) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Todoes is not required by pop and may be deleted
type Todos []Todo

// String is not required by pop and may be deleted
func (t Todos) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Todo) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Todo) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Todo) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
