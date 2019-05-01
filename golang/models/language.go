package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type Language struct {
	ID            uint64    `json:"id" db:"id"`
	Name          string    `json:"name,omitempty" db:"name"`
	Description   string    `json:"description,omitempty" db:"description"`
	Documentation string    `json:"documentation,omitempty" db:"documentation"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (l Language) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Languages is not required by pop and may be deleted
type Languages []Language

// String is not required by pop and may be deleted
func (l Languages) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Language) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Language) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Language) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
