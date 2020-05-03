package models

import (
	"encoding/json"
	"fmt"
	"management/utils"
	"strings"
	"time"
	"unicode"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   int64      `json:"id" db:"id"`
	Email                string     `json:"email" form:"email" db:"email"`
	Name                 string     `json:"name" db:"name"`
	Surname              string     `json:"surname" db:"surname"`
	Password             string     `json:"-" db:"password"`
	RememberToken        string     `json:"remember_token" db:"remember_token"`
	Slug                 string     `json:"slug" db:"slug"`
	Type                 int        `json:"type" db:"type"`
	JoinedAt             nulls.Time `json:"joined_at" db:"joined_at"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`
	PasswordPlain        string     `json:"password,omitempty"  form:"password" db:"-"`
	PasswordConfirmation string     `json:"password_confirmation,omitempty" form:"password_confirmation"  db:"-" `
	//Relationships
	Settings  UserSetting `has_one:"user_settings"`
	Tasks     Tasks       `many_to_many:"users_tasks"`
	Projects  Projects    `many_to_many:"users_projects"`
	Languages Languages   `many_to_many:"users_languages"`
	Comments  Comments    `has_many:"comments" order_by:"created_at desc"`
	Accounts  Platforms   `many_to_many:"users_platforms"`
	Actions   UserActions `json:"user_actions,omitmepty" has_many:"user_actions"`
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
		&validators.EmailIsPresent{Field: u.Email, Name: "Email"},
		&validators.FuncValidator{
			Field:   "Name",
			Name:    "Name",
			Message: "%s minimum length is 3",
			Fn: func() bool {
				return len(u.Name) > 3
			}},
		&validators.FuncValidator{
			Field:   "Surname",
			Name:    "Surname",
			Message: "%s minimum length is 3",
			Fn: func() bool {
				return len(u.Surname) > 3
			}},
		&validators.FuncValidator{
			Field:   "Password",
			Name:    "Password",
			Message: "%s is not strong enough",
			Fn: func() bool {
				if len(u.PasswordPlain) < 6 {
					return false
				}
				var number, upper, special bool
				letters := 0
				for _, c := range u.PasswordPlain {
					switch {
					case unicode.IsNumber(c):
						number = true
					case unicode.IsUpper(c):
						upper = true
						letters++
					case unicode.IsPunct(c) || unicode.IsSymbol(c):
						special = true
					case unicode.IsLetter(c) || c == ' ':
						letters++
					default:
						//return false, false, false, false
					}
				}

				return number && upper && special
			},
		},
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
		&validators.StringsMatch{Name: "PasswordConfirmation", Field: u.PasswordPlain, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) create(tx *pop.Connection) error {
	if u.Email != "" {
		u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	}

	if u.PasswordPlain != "" && u.PasswordConfirmation == u.PasswordPlain {

		ph, err := bcrypt.GenerateFromPassword([]byte(u.PasswordPlain), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		u.Password = string(ph)

		u.PasswordConfirmation = ""
		u.PasswordPlain = ""
	}

	return nil
}

func (u *User) BeforeCreate(tx *pop.Connection) error {

	const (
		typeAction = "confirm_email"
	)

	token, err := utils.GenerateToken(fmt.Sprintf("%d_%s_%s", time.Now().Unix(), typeAction, u.Email))

	if err != nil {
		return err
	}

	u.Actions = append(u.Actions, UserAction{
		Type:  typeAction,
		Token: token,
	})

	if u.Name != "" && u.Surname != "" {
		err := u.GenerateSlug()

		if err != nil {
			return err
		}
	}

	return u.create(tx)
}

func (u *User) BeforeUpdate(tx *pop.Connection) error {

	return u.create(tx)
}

func (u *User) GenerateSlug() error {
	userSlug := slug.Make(fmt.Sprintf("%s_%s", u.Name, u.Surname))

	count, err := DB.Where("slug LIKE '%' || ? || '%' ", userSlug).Count(&User{})

	if err != nil {

		return err
	}

	if count > 0 {
		userSlug = fmt.Sprintf("%s_%d", userSlug, count)
	}

	u.Slug = userSlug

	return nil
}

func (u *User) CacheKey() string {
	return u.ToCacheKey(u.ID)
}

func (u *User) ToCacheKey(id int64) string {
	return fmt.Sprintf("USER:%d", id)
}
