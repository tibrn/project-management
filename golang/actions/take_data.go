package actions

import (
	"context"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/myWebsite/golang/models"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Repository struct {
	DatabaseID  githubv4.Int
	Description githubv4.String
	HomepageURL githubv4.URI
	IsPrivate   githubv4.Boolean
	LicenseInfo License
	Name        githubv4.String
}

type License struct {
	Body        githubv4.String
	Description githubv4.String
	Name        githubv4.String
	Nickname    githubv4.String
	Key         githubv4.String
	URL         githubv4.URI
}

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
	variables := map[string]interface{}{
		"login":              githubv4.String(userPlatform.Username),
		"repositoriesCursor": (*githubv4.String)(nil),
	}
	var q struct {
		User struct {
			ID           githubv4.ID
			Repositories struct {
				Nodes    []Repository
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"repositories(first: 15, after: $repositoriesCursor)"`
		} `graphql:"user(login: \"BTiberiu99\")"`
	}
	// Get comments from all pages.
	var allRepositories []Repository
	for {
		err := client.Query(context.Background(), &q, variables)
		if err != nil {
			return err
		}
		allRepositories = append(allRepositories, q.User.Repositories.Nodes...)
		if !q.User.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubv4.NewString(q.User.Repositories.PageInfo.EndCursor)
	}

	err = client.Query(context.Background(), &q, variables)
	if err != nil {
		// Handle error.
		fmt.Println(err)
	} else {
		fmt.Println(&q)
	}

	// fmt.Println(q)
	return nil
}

func projectLanguages(PlatformID string, ID int64) error {
	return nil
}

func projectLicense(PlatformID string, ID int64) error {
	return nil
}
