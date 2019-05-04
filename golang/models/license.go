package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type License struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name,omitempty" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	Body        string    `json:"body,omitempty" db:"body"`
	Nickname    string    `json:"nickname,omitempty" db:"nickname"`
	Key         string    `json:"key,omitempty" db:"key"`
	URL         string    `json:"url,omitempty" db:"url"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	//Relationships
	Project *Project `has_one:"projects" db:"-"`
}

// String is not required by pop and may be deleted
func (l License) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Licenses is not required by pop and may be deleted
type Licenses []License

// String is not required by pop and may be deleted
func (l Licenses) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *License) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *License) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *License) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
