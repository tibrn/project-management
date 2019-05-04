package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type ProjectLanguage struct {
	ID         int64     `json:"id" db:"id"`
	ProjectID  uuid.UUID `json:"project_id" db:"project_id"`
	LanguageID int64     `json:"language_id" db:"language_id"`
	Primary    bool      `json:"primary" db:"is_primary"`
	Usage      float64   `json:"usage" db:"usage"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p ProjectLanguage) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// ProjectLanguages is not required by pop and may be deleted
type ProjectLanguages []ProjectLanguage

// String is not required by pop and may be deleted
func (p ProjectLanguages) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *ProjectLanguage) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *ProjectLanguage) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *ProjectLanguage) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (ProjectLanguage) TableName() string {
	return "projects_languages"
}
