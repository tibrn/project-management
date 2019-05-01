package actions

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/myWebsite/golang/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(Error{Message: T.Translate(c, "message.server.error")}))
	}
	tx := c.Value("tx").(*pop.Connection)

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(u.Email)).First(u)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return c.Render(http.StatusForbidden, r.JSON(Error{Message: T.Translate(c, "user.not.found")}))
		}
		return c.Render(http.StatusInternalServerError, r.JSON(Error{Message: T.Translate(c, "message.server.error")}))
	}
	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.PasswordPlain))

	if err != nil {
		return c.Render(http.StatusForbidden, r.JSON(Error{Message: T.Translate(c, "user.not.found")}))
	}

	c.Session().Set("current_user_id", u.ID)
	tx.Load(u, "Settings")
	return c.Render(http.StatusOK, r.JSON(MessageData{
		Message:     T.Translate(c, "user.succes.login"),
		MessageType: "succes",
		Data:        u,
	}))
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	return c.Render(http.StatusOK, r.JSON(Success{Message: T.Translate(c, "user.succes.logout")}))
}

// AuthRefresh sends refreshed data to browser
func AuthRefresh(c buffalo.Context) error {
	user, ok := c.Value("current_user").(*models.User)

	if !ok {
		return c.Render(http.StatusForbidden, r.JSON(Error{Message: T.Translate(c, "user.not.found")}))
	}

	tx, ok := c.Value("tx").(*pop.Connection)

	tx.Load(user, "Settings")

	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(Error{Message: T.Translate(c, "message.server.error")}))
	}

	return c.Render(http.StatusOK, r.JSON(user))
}
