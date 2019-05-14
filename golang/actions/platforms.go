package actions

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/events"
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

	exist := &models.UserPlatform{}
	err = models.DB.Where("username = ? OR id_on_platform = ? ", userGithub.GetLogin(), userGithub.GetID()).Select("id", "username").First(exist)

	//S-a gasit un user
	if err == nil {
		c.Session().Set("exist", map[string]string{
			"name":     exist.Username,
			"platform": "Github",
		})
		return c.Redirect(http.StatusMovedPermanently, "/welcome")
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
	userPlatform.Username = userGithub.GetLogin()
	userPlatform.URL = userGithub.GetURL()

	verrs, err := models.DB.ValidateAndCreate(userPlatform)

	if err != nil {
		return HTTP500(c)
	}

	if verrs.HasAny() {
		return HTTP403(c, T.Translate(c, "token.not.created"), verrs.Errors)
	}

	if user.Settings.Avatar == "" {
		settings := user.Settings
		settings.Avatar = userGithub.GetAvatarURL()
		models.DB.ValidateAndUpdate(&settings)
	}

	e := events.Event{
		Kind:    "platform:projects_github",
		Payload: events.Payload{"user_id": user.ID},
	}

	if err := events.Emit(e); err != nil {
		return HTTP500(c)
	}
	c.Session().Set("name", userPlatform.Username)
	return c.Redirect(http.StatusMovedPermanently, "/welcome")
}

func (Platform) GithubRedirect(c buffalo.Context) error {

	fmt.Println("CEVA")
	URL := envy.Get("GITHUB_AUTHORIZE_URL", "https://github.com/login/oauth/authorize")
	URL += "?client_id=" + envy.Get("GITHUB_CLIENT_ID", "")
	URL += "&redirect_uri=" + envy.Get("APP_URL", "") + "/github/callback"

	fmt.Println("GITHUB", URL)
	return c.Redirect(http.StatusMovedPermanently, URL)
}
