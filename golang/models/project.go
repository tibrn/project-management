package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type Project struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	PlatformID   int64        `json:"platform_id" db:"platform_id"`
	LicenseID    nulls.Int64  `json:"license_id" db:"license_id"`
	IDOnPlatform int64        `json:"id_on_platform" db:"id_on_platform"`
	Name         string       `json:"name,omitempty" db:"name"`
	Description  nulls.String `json:"description,omitempty" db:"description"`
	URL          string       `json:"url" db:"url"`
	CreatedAt    time.Time    `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at,omitempty" db:"updated_at"`
	//Relationships
	Tasks     Tasks     `json:"tasks,omitempty" many_to_many:"projects_tasks" db:"-"`
	Languages Languages `json:"languages,omitempty" many_to_many:"projects_languages" db:"-"`
	License   *License  `json:"licenses,omitempty" belongs_to:"licenses" db:"-"`
	Platform  Platform  `json:"platform,omitempty" belongs_to:"platforms" db:"-"`
	Users     Users     `json:"users,omitempty" many_to_many:"users_projects" db:"-"`
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
	return validate.NewErrors(), nil
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
