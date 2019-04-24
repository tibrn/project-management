package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/myWebsite/golang/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
