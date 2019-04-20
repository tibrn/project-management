package actions

import (
	"github.com/gobuffalo/buffalo"
)

type Platform struct{}

func (Platform) GithubCallback(c buffalo.Context) {

}
