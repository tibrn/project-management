package actions

import (
	"context"
	"fmt"
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

	client := github.NewClient(oauthConfig.Client(oauth2.NoContext, tkn))

	userGithub, _, err := client.Users.Get(context.Background(), "")
	fmt.Println(userGithub)
	exist := &models.UserPlatform{}
	err = models.DB.Where("username = ? OR id_on_platform = ? ", userGithub.GetName(), userGithub.GetID()).Select("id").First(exist)

	//S-a gasit un user
	if err == nil {
		return HTTP403(c, T.Translate(c, "user.platform.exist"))
	}

	limits, _, err := client.RateLimits(context.Background())

	user := c.Value("current_user").(*models.User)

	userPlatform := &models.UserPlatform{}
	userPlatform.UserID = user.ID
	userPlatform.PlatformID = 1
	userPlatform.IDOnPlatform = userGithub.GetID()
	userPlatform.Token = tkn.AccessToken
	userPlatform.TokenType = tkn.TokenType
	userPlatform.Limit = int64(limits.Core.Limit)
	userPlatform.ResetAt = nulls.NewTime(limits.Core.Reset.Time)
	userPlatform.Username = userGithub.GetName()
	userPlatform.URL = userGithub.GetLogin()

	verrs, err := models.DB.ValidateAndCreate(userPlatform)

	if err != nil {
		return HTTP500(c)
	}

	if verrs.HasAny() {
		return HTTP403(c, T.Translate(c, "token.not.created"), verrs.Errors)
	}

	if user.Settings.Avatar == "" {
		user.Settings.Avatar = userGithub.GetAvatarURL()
		fmt.Println(user.Settings)
		models.DB.ValidateAndUpdate(user.Settings)
	}
	return c.Redirect(http.StatusMovedPermanently, "/welcome")
}
