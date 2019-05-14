package actions

import (
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/myWebsite/golang/models"
)

func ActivateEmail(c buffalo.Context) error {
	// Allocate an empty User
	verify := &models.UserVerify{}
	// Bind user to the html form elements
	if err := c.Bind(verify); err != nil {
		return HTTP500(c)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return HTTP500(c)
	}

	err := tx.Where("token = ?", verify.Token).First(verify)

	if err != nil {
		return HTTP500(c)
	}

	user := &models.User{
		ID:       verify.UserID,
		JoinedAt: nulls.NewTime(time.Now()),
	}

	err = tx.Update(user, "type", "slug", "remember_token", "email", "name", "password")

	if err != nil {
		return HTTP500(c)
	}

	if uid := c.Session().Get("current_user_id"); uid == nil {
		c.Session().Set("current_user_id", user.ID)
	}

	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}
