package actions

import "github.com/gobuffalo/buffalo"

func ProjectsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("projects/index.html"))
}
