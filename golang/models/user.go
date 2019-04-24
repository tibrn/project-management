package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   uint64    `json:"id" db:"id"`
	Email                string    `json:"email" db:"email"`
	Password             string    `json:"password" db:"password"`
	RememberToken        string    `json:"remember_token" db:"remember_token"`
	Slug                 string    `json:"slug" db:"slug"`
	Type                 int8      `json:"type" db:"type"`
	JoinedAt             time.Time `json:"joined_at" db:"joined_at"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	PasswordPlain        string    `json:"-" db:"-" form:"password_plain"`
	PasswordConfirmation string    `json:"-" db:"-" form:"password_confirmation"`
	//Relationships
	Settings  UserSettings `has_one:"user_settings"`
	Tasks     Tasks        `many_to_many:"users_tasks"`
	Projects  Projects     `many_to_many:"users_projects"`
	Languages Languages    `many_to_many:"users_languages"`
	Comments  Comments     `has_many:"comments" order_by:"created_at desc"`
	Accounts  Platforms    `many_to_many:"users_platforms"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		// check to see if the email address is already taken:
		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "Email",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("email = ?", u.Email)
				if u.ID != 0 {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.PasswordPlain, Name: "PasswordPlain"},
		&validators.StringIsPresent{Field: u.PasswordConfirmation, Name: "PasswordConfirmation"},
		&validators.StringsMatch{Name: "PasswordPlain", Field: u.PasswordPlain, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	ph, err := bcrypt.GenerateFromPassword([]byte(u.PasswordPlain), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.Password = string(ph)
	return tx.ValidateAndCreate(u)
}
