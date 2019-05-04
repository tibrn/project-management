package actions

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/myWebsite/golang/models"
	"github.com/pkg/errors"
)

func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about/index.html"))
}

func ProjectsHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	projects := &models.Projects{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// authID := c.Value("current_user_id").(int64)

	if err := q.All(projects); err != nil {
		return errors.WithStack(err)
	}

	c.Set("pagination", q.Paginator)
	c.Set("projects", projects)

	return c.Render(200, r.HTML("projects/index.html"))
}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("home/index.html"))
}

// RedirectPlatformHandler is handler to projects page
func RedirectPlatformHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("platform/index.html"))
}

// VueHandler is a handler that server Vue Application
func VueHandler(c buffalo.Context) error {

	c.Set("locale", strings.Split(strings.Split(c.Request().Header.Get("Accept-Language"), ";")[0], ",")[0])
	// Get the DB connection from the context
	user, _ := c.Value("current_user").(*pop.Connection)
	if user != nil {
		// Get the DB connection from the context
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return errors.WithStack(errors.New("no transaction found"))
		}
		tx.Load(user, "Settings")

		c.Set("current_user", user)
	}

	return c.Render(200, r.HTML("vue/index.html", "layout/empty.html"))
}

//Test is used to make test to retrive data from platforms
func Test(c buffalo.Context) error {
	ID := c.Session().Get("current_user_id").(int64)
	Init(ID)
	return nil
}
