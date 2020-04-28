package actions

import (
	"management/enums"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func InternalError(c buffalo.Context) error {
	return c.Render(http.StatusInternalServerError, r.JSON(Response{
		Message: T.Translate(c, "errors.internal"),
		Type:    enums.Error,
	}))
}
