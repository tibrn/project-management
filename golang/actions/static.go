package actions

import (
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

	user := &models.User{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params()).Eager("Projects")

	currentUser := c.Value("current_user").(*models.User)
	// Retrieve all Projects from the DB
	if err := q.Find(user, currentUser.ID); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("projects", user.Projects)
	return c.Render(200, r.HTML("projects/index.html"))
}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("home/index.html"))
}
