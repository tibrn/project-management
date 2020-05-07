package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Project struct {
	ID          uuid.UUID `json:"id" db:"id"`
	PlatformID  int64     `json:"platform_id" db:"platform_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	//Relationships
	Tasks     Tasks     `json:"tasks,omitempty" many_to_many:"projects_tasks"`
	Languages Languages `json:"languages,omitempty" many_to_many:"projects_languages"`
	License   *License  `json:"license,omitempty" has_one:"license"`
}

// String is not required by pop and may be deleted
func (p Project) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Projects is not required by pop and may be deleted
type Projects []Project

// String is not required by pop and may be deleted
func (p Projects) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Project) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(&validators.FuncValidator{
		Field:   "Name",
		Name:    "Name",
		Message: "%s minimum length is 4",
		Fn: func() bool {
			return len(p.Name) >= 4
		}},
		&validators.FuncValidator{
			Field:   "Description",
			Name:    "Description",
			Message: "%s minimum length is 100",
			Fn: func() bool {
				return len(p.Description) >= 100
			}}, &validators.FuncValidator{
			Field:   "Description",
			Name:    "Description",
			Message: "%s maximum length is 255",
			Fn: func() bool {
				return len(p.Description) <= 255
			}}), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
