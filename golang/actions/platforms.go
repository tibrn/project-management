package actions

import (
	"context"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
	"github.com/google/go-github/github"
	"github.com/myWebsite/golang/models"
	"golang.org/x/oauth2"
)

type Platform struct{}

func (Platform) GithubCallback(c buffalo.Context) error {

	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("GITHUB_AUTHORIZE_URL"),
			TokenURL: os.Getenv("GITHUB_TOKEN_URL"),
		},
	}

	tkn, err := oauthConfig.Exchange(oauth2.NoContext, c.Param("code"))
	if err != nil {
		// Handle Error
		return HTTP500(c)
	}

	if !tkn.Valid() {
		return HTTP403(c, T.Translate(c, "message.token.invalid"))
	}
	tx := models.DB

	client := github.NewClient(oauthConfig.Client(oauth2.NoContext, tkn))

	limits, _, err := client.RateLimits(context.Background())

	user := c.Value("current_user").(*models.User)

	userPlatform := &models.UserPlatform{}
	userPlatform.UserID = user.ID
	userPlatform.PlatformID = 1
	userPlatform.Token = tkn.AccessToken
	userPlatform.TokenType = tkn.TokenType
	userPlatform.Limit = int64(limits.Core.Limit)
	userPlatform.ResetAt = nulls.NewTime(limits.Core.Reset.Time)

	verrs, err := tx.ValidateAndCreate(userPlatform)

	if err != nil {
		return HTTP500(c)
	}

	if verrs.HasAny() {
		return HTTP403(c, T.Translate(c, "token.not.created"), verrs.Errors)
	}
	return c.Redirect(http.StatusMovedPermanently, "/welcome")
}
