package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type UserPlatform struct {
	ID         int64      `json:"id" db:"id"`
	UserID     int64      `json:"user_id" db:"user_id"`
	PlatformID int64      `json:"platform_id" db:"platform_id"`
	Token      string     `json:"-" db:"token"`
	TokenType  string     `json:"-" db:"token_type"`
	Limit      int64      `json:"limit,omitempty" db:"limit_requests"`
	ResetAt    nulls.Time `json:"reset_at,omitempty" db:"reset_at"`
	CreatedAt  time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	User       User       `json:"user,omitempty" belongs_to:"users" db:"-"`
	Platform   Platform   `json:"platform,omitempty" belongs_to:"platforms" db:"-"`
}

// String is not required by pop and may be deleted
func (u UserPlatform) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserPlatforms is not required by pop and may be deleted
type UserPlatforms []UserPlatform

// String is not required by pop and may be deleted
func (u UserPlatforms) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserPlatform) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserPlatform) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserPlatform) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// TableName overrides the table name used by Pop.
func (u UserPlatform) TableName() string {
	return "users_platforms"
}
