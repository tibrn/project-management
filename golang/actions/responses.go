package actions

import (
	"management/enums"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func InternalError(c buffalo.Context) error {
	return Error(c, http.StatusInternalServerError, enums.ErrorsInternal)
}

func Error(c buffalo.Context, status int, tranlsate string, errs ...interface{}) error {
	var d interface{}

	if len(errs) > 0 {
		d = errs[0]
	}
	return c.Render(status, r.JSON(Response{
		Message: T.Translate(c, tranlsate),
		Type:    enums.Error,
		Errors:  d,
	}))
}

func Success(c buffalo.Context, tranlsate string, data ...interface{}) error {

	var d interface{}

	if len(data) > 0 {
		d = data[0]
	}

	return c.Render(http.StatusOK, r.JSON(Response{
		Message: T.Translate(c, tranlsate),
		Type:    enums.Success,
		Data:    d,
	}))
}
