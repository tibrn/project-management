package actions

import (
	"fmt"
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

	projects := []*models.Project{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// authID := c.Value("current_user_id").(int64)

	if err := q.All(&projects); err != nil {
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

	exist := c.Session().Get("exist")

	name := c.Session().Get("name")

	if exist != nil {
		c.Set("exist", exist)
		c.Session().Delete("exist")
	} else if name != nil {
		c.Set("name", name)
		c.Session().Delete("name")
	}
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
	// tx, ok := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return errors.WithStack(errors.New("no transaction found"))
	// }
	test := &models.Languages{}
	// tx.Where("user_id = ", 1).All(test)
	err := models.DB.RawQuery("SELECT * FROM languages WHERE id IN (SELECT language_id FROM projects_languages WHERE project_id IN (SELECT project_id FROM users_projects WHERE user_id =1 ) )").All(test)

	fmt.Println(err)
	fmt.Println(test)
	return nil
}
