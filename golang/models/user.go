package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	ADMIN  = 1
	CLIENT = 2
)

type User struct {
	ID                   int64        `json:"id" db:"id"`
	Name                 string       `json:"name,omitempty" form:"name" db:"name" `
	Email                string       `json:"email,omitempty" form:"email" db:"email" `
	Password             string       `json:"-" db:"password"`
	RememberToken        nulls.String `json:"-" db:"remember_token"`
	Slug                 nulls.String `json:"slug,omitempty" db:"slug"`
	Type                 int8         `json:"type,omitempty" db:"type"`
	JoinedAt             nulls.Time   `json:"joined_at,omitempty" db:"joined_at"`
	CreatedAt            time.Time    `json:"-" db:"created_at"`
	UpdatedAt            time.Time    `json:"-" db:"updated_at"`
	PasswordPlain        string       `json:"password,omitempty" form:"password" db:"-" `
	PasswordConfirmation string       `json:"password_confirmation,omitempty"  form:"password_confirmation" db:"-"`
	//Relationships
	Settings  UserSetting `json:"settings,omitempty" has_one:"user_setting" db:"-"`
	Tasks     Tasks       `json:"tasks,omitempty" many_to_many:"users_tasks" db:"-"`
	Projects  Projects    `json:"projects,omitempty" many_to_many:"users_projects" db:"-"`
	Languages Languages   `json:"languages,omitempty" many_to_many:"users_languages" db:"-"`
	Comments  Comments    `json:"comments,omitempty" has_many:"comments" db:"-" order_by:"created_at desc"`
	Accounts  Platforms   `json:"accounts,omitempty" many_to_many:"users_platforms" db:"-"`
}

//SafeUser is used to send user infromation
//back to client without private information
type SafeUser User

// String is not required by pop and may be deleted
func (u User) String() string {
	u.PasswordConfirmation = ""
	u.PasswordPlain = ""
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u User) MarshalJSON() ([]byte, error) {
	user := SafeUser(u)
	user.PasswordConfirmation = ""
	user.PasswordPlain = ""
	return json.Marshal((*SafeUser)(&user))
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
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
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
	u.Slug = u.unqiueSlug()

	fmt.Println(u)
	return tx.Eager().ValidateAndCreate(u)
}

func (u *User) isAdmin() bool {
	return u.Type == ADMIN
}

func (u *User) isClient() bool {
	return u.Type == CLIENT
}

func (u *User) unqiueSlug() nulls.String {
	slug := slug.Make(u.Name)
	nr, _ := DB.Where("slug LIKE '%' || ? || '%' ", slug).Select("id").Count(u)
	if nr > 0 {
		slug += fmt.Sprintf("%d", nr+1)
	}
	return nulls.NewString(slug)
}
