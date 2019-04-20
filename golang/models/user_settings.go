package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type UserSetting struct {
	ID        uint64    `json:"id" db:"id"`
	UserID    uint64    `json:"user_id" db:"user_id"`
	Avatar    string    `json:"avatar" db:"avatar"`
	User      *User     `belongs_to:"user"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u UserSetting) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserSettings is not required by pop and may be deleted
type UserSettings []UserSetting

// String is not required by pop and may be deleted
func (u UserSettings) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserSetting) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserSetting) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserSetting) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
