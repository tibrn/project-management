package actions

import (
	"context"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/myWebsite/golang/models"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func githubProjects(c buffalo.Context) error {
	ID := c.Session().Get("current_user_id").(int64)

	// Get the DB connection from the context
	userPlatform := &models.UserPlatform{}
	err := models.DB.Where("user_id = ?", ID).First(userPlatform)
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: userPlatform.Token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var q struct {
		Viewer struct {
			Login          githubv4.String
			CreatedAt      githubv4.DateTime
			IsBountyHunter githubv4.Boolean
			BioHTML        githubv4.HTML
			WebsiteURL     githubv4.URI
		}
	}

	err = client.Query(context.Background(), &q, nil)
	if err != nil {
		// Handle error.
	}

	fmt.Println(q)
	return nil
}
