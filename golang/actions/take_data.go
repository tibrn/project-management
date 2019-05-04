package actions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"github.com/myWebsite/golang/models"
	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const (
	KB = 1024
	MB = KB * 1024
)

type Data struct {
	User     models.UserPlatform
	Projects []Project
	client   *githubv4.Client
}

type Language struct {
	Color string
	Name  string
	Size  int64
}

type Repository struct {
	DatabaseID  int64
	Description string
	URL         string
	IsPrivate   bool
	LicenseInfo License
	Name        string
	//Useful data
	PrimaryLanguage struct {
		Name string
	}
}

type License struct {
	Body        string
	Description string
	Name        string
	Nickname    string
	Key         string
	URL         string
}

type RelationshipLanguage struct {
	Model models.Language
	Usage float64
}

type Project struct {
	DatabaseID  int64
	Description string
	URL         string
	Private     bool
	// LicenseInfo License
	Name string

	//Useful data
	PrimaryLanguage string
	Languages       []Language
	License         License
}

type queryRespositories struct {
	User struct {
		Repositories struct {
			Nodes    []Repository
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"repositories(first: 50, after: $repositoriesCursor, isFork: false, affiliations : [OWNER])"`
	} `graphql:"user(login: $login)"`
}

type queryLanguages struct {
	Repository struct {
		Languages struct {
			Edges []struct {
				Node struct {
					Color string
					Name  string
				}
				Size int64
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"languages(first: 15, after: $languagesCursor)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

//Init start to take all accounts from platforms of an user
//intialize http client with token for each
//and start to take projects from platforms
func Init(ID int64) error {
	//TAKE USER_PLATFORM FROM DB
	userPlatforms := []models.UserPlatform{}
	err := models.DB.Where("user_id = ? AND COALESCE(last_updated_at,DATE '0001-01-01') < ?", ID, time.Now().Add(-24*time.Hour)).All(&userPlatforms)
	if err != nil {

		fmt.Println("TAKE USER_PLATFORM", err)
		return err
	}
	dataPlatforms := make([]Data, len(userPlatforms))
	for i := range userPlatforms {
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: userPlatforms[i].Token},
		)
		httpClient := oauth2.NewClient(context.Background(), src)
		client := githubv4.NewClient(httpClient)
		//SET CLIENT
		dataPlatforms[i].client = client

		//SET USER
		dataPlatforms[i].User = userPlatforms[i]

		respositories, err := dataPlatforms[i].takeRepositories()
		if err != nil {
			fmt.Println(err)
		}

		dataPlatforms[i].transfRepositories(respositories)
		if len(dataPlatforms[i].Projects) > 0 {
			go func() {
				dataPlatforms[i].update()
			}()
		}

	}

	return nil
}

func (d *Data) takeRepositories() (repositories []Repository, err error) {
	variables := map[string]interface{}{
		"repositoriesCursor": (*githubv4.String)(nil),
		"login":              githubv4.String(d.User.Username),
	}
	qReps := queryRespositories{}
	for {
		err = d.client.Query(context.Background(), &qReps, variables)
		if err != nil {
			return repositories, err
		}
		repositories = append(repositories, qReps.User.Repositories.Nodes...)
		if !qReps.User.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubv4.NewString(qReps.User.Repositories.PageInfo.EndCursor)
	}
	return repositories, nil
}

func (d *Data) takeLanguages(name string) (languages []Language, err error) {
	variables := map[string]interface{}{
		"languagesCursor": (*githubv4.String)(nil),
		"owner":           githubv4.String(d.User.Username),
		"name":            githubv4.String(name),
	}

	qLangs := queryLanguages{}
	for {
		err = d.client.Query(context.Background(), &qLangs, variables)
		if err != nil {
			fmt.Println("TAKE LANGUAGES", err)
			return languages, err
		}
		for i := range qLangs.Repository.Languages.Edges {
			languages = append(languages, Language{
				Name:  qLangs.Repository.Languages.Edges[i].Node.Name,
				Color: qLangs.Repository.Languages.Edges[i].Node.Color,
				Size:  qLangs.Repository.Languages.Edges[i].Size,
			})
		}

		if !qLangs.Repository.Languages.PageInfo.HasNextPage {
			break
		}
		variables["languagesCursor"] = githubv4.NewString(qLangs.Repository.Languages.PageInfo.EndCursor)
	}

	return languages, nil
}

func (d *Data) transfRepositories(repositories []Repository) error {
	if len(repositories) < 1 {
		return errors.New("Nu s-a putut prelua niciun repository")
	}

	d.Projects = make([]Project, len(repositories))
	fmt.Println("TRANSF REPOSITORIES")
	for i := range repositories {
		d.Projects[i].URL = repositories[i].URL
		d.Projects[i].Name = repositories[i].Name
		d.Projects[i].Private = repositories[i].IsPrivate
		d.Projects[i].License = repositories[i].LicenseInfo
		d.Projects[i].Description = repositories[i].Description
		d.Projects[i].DatabaseID = repositories[i].DatabaseID
		d.Projects[i].PrimaryLanguage = repositories[i].PrimaryLanguage.Name

		languages, err := d.takeLanguages(repositories[i].Name)
		if err != nil {
			fmt.Println(err)
		}
		d.Projects[i].Languages = languages
	}
	return nil
}

func (d *Data) update() {
	for i := range d.Projects {
		d.updateProject(&d.Projects[i])
	}

	fmt.Println("PROJECTS UPDATED FOR USER WITH ID =", d.User.ID)
}

func (d *Data) updateProject(data *Project) {
	project := &models.Project{}

	err := models.DB.Transaction(func(tx *pop.Connection) error {

		err := tx.Where("id_on_platform = ?", data.DatabaseID).First(project)

		license := d.FindCreateLicense(&data.License)

		languages := d.FindCreateLanguages(data.Languages)

		project.Name = data.Name
		project.URL = data.URL
		project.Description = nulls.NewString(data.Description)

		query := tx

		if license != nil {
			fmt.Println(license)
			project.License = *license
			query = query.Eager("License")
		}

		var verrs *validate.Errors
		if errors.Cause(err) == sql.ErrNoRows {
			project.PlatformID = 1
			project.IDOnPlatform = data.DatabaseID
			verrs, err = query.ValidateAndCreate(project)
		} else {
			verrs, err = query.ValidateAndUpdate(project)
		}

		if err != nil {
			fmt.Println("PROJECT GITHUB", err)
			return err
		}

		if verrs.HasAny() {
			fmt.Println("PROJECT GITHUB", verrs.Errors)
			return errors.New("Couldn't create project")
		}

		if err == nil {
			d.User.LastUpdatedAt = nulls.NewTime(time.Now())
			tx.ValidateAndUpdate(&d.User)

			userProject := &models.UserProject{}

			err := tx.Where("project_id = ? AND user_id = ?", project.ID, d.User.UserID).First(userProject)

			if errors.Cause(err) == sql.ErrNoRows {
				userProject.UserID = d.User.UserID
				userProject.ProjectID = project.ID

				verrs, err := tx.ValidateAndCreate(userProject)

				fmt.Println("PROJEC CREATED")
				if err != nil {
					fmt.Println("CREATE RELATIONSHIP ", err)
					return err
				}

				if verrs.HasAny() {
					fmt.Println("CREATE RELATIONSHIP ", verrs.Errors)
					return errors.New("Couldn't create realtionship for project and user")
				}
			}

			err = d.UpdateCreateUserLanguagesRelationships(languages, tx)

			if err != nil {
				return err
			}

			err = d.UpdateCreateProjectLanguagesRelationships(project.ID, data.PrimaryLanguage, languages, tx)

			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("TRANSACTION GITHUB", err)
	}
}

func (d *Data) UpdateCreateUserLanguagesRelationships(languages []RelationshipLanguage, tx *pop.Connection) error {

	for i := range languages {
		relation := &models.UserLanguage{}
		err := tx.Where("language_id = ? AND user_id = ?", languages[i].Model.ID, d.User.UserID).First(relation)

		relation.Proficiency = languages[i].Usage / KB / 100
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				relation.UserID = d.User.UserID
				relation.LanguageID = languages[i].Model.ID

				verrs, err := tx.ValidateAndCreate(relation)

				if err != nil {
					fmt.Println("CREATE REALTIONSHIP USER LANGUAGE ", err)
					return err
				}

				if verrs.HasAny() {
					fmt.Println("CREATE REALTIONSHIP USER LANGUAGE ", verrs.Errors)
					return errors.New("Couldn't create realtionship for language and user")
				}
			} else {
				fmt.Println("REALTIONSHIP USER LANGUAGE", err)
				return err
			}
		} else {
			verrs, err := tx.ValidateAndUpdate(relation)
			if verrs.HasAny() {
				fmt.Println("UPDATE REALTIONSHIP USER LANGUAGE", verrs.Errors)
			}

			if err != nil {
				fmt.Println("UPDATE REALTIONSHIP USER LANGUAGE", err)
			}
		}

	}
	return nil
}

func (d *Data) UpdateCreateProjectLanguagesRelationships(ID uuid.UUID, primary string, languages []RelationshipLanguage, tx *pop.Connection) error {

	for i := range languages {
		relation := &models.ProjectLanguage{}
		err := tx.Where("language_id = ? AND project_id = ?", languages[i].Model.ID, ID).First(relation)

		relation.Usage = languages[i].Usage
		relation.Primary = languages[i].Model.Name == primary
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				relation.ProjectID = ID
				relation.LanguageID = languages[i].Model.ID

				verrs, err := tx.ValidateAndCreate(relation)

				if err != nil {
					fmt.Println("CREATE REALTIONSHIP PROJECT LANGUAGE ", err)
					return err
				}

				if verrs.HasAny() {
					fmt.Println("CREATE REALTIONSHIP PROJECT LANGUAGE ", verrs.Errors)
					return errors.New("Couldn't create realtionship for project and language")
				}
			} else {
				fmt.Println("REALTIONSHIP PROJECT LANGUAGE", err)
				return err
			}
		} else {
			verrs, err := tx.ValidateAndUpdate(relation)
			if verrs.HasAny() {
				fmt.Println("UPDATE REALTIONSHIP PROJECT LANGUAGE", verrs.Errors)
			}

			if err != nil {
				fmt.Println("UPDATE REALTIONSHIP PROJECT LANGUAGE", err)
			}
		}

	}
	return nil
}

//FindCreateLicense ... Search for or license and retrive
//or create license
func (d *Data) FindCreateLicense(data *License) *models.License {

	if data.Name == "" || data.Nickname == "" {
		return nil
	}
	license := &models.License{}

	err := models.DB.Where("nickname = ? OR name = ?", data.Name, data.Nickname).First(license)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			license.Key = data.Key
			license.URL = data.URL
			license.Body = data.Body
			license.Nickname = data.Nickname
			license.Description = data.Description

			verrs, err := models.DB.ValidateAndCreate(license)

			if err != nil {
				fmt.Println("CREATE LICENSE", err)
				return nil
			}

			if verrs.HasAny() {
				fmt.Println("CREATE LICENSE", verrs.Errors)
				return nil
			}
		} else {
			return nil
		}
	}

	return license
}

func (d *Data) FindCreateLanguages(data []Language) []RelationshipLanguage {
	languages := make([]RelationshipLanguage, len(data))

	for i := range data {
		err := models.DB.Where("name = ?", data[i].Name).First(&languages[i].Model)

		if errors.Cause(err) == sql.ErrNoRows {
			languages[i].Model.Color = data[i].Color
			languages[i].Model.Name = data[i].Name

			verrs, err := models.DB.ValidateAndCreate(&languages[i].Model)

			if err != nil {
				fmt.Println("CREATE LANGUAGES", err)
			}

			if verrs.HasAny() {
				fmt.Println("CREATE LANGUAGE", verrs.Errors)
			}

			languages[i].Usage = float64(data[i].Size)
		}
	}
	return languages
}

// func githubProjects(c buffalo.Context) error {
// 	ID := c.Session().Get("current_user_id").(int64)

// 	// Get the DB connection from the context
// 	userPlatform := &models.UserPlatform{}
// 	err := models.DB.Where("user_id = ?", ID).First(userPlatform)
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: userPlatform.Token},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	client := githubv4.NewClient(httpClient)
// 	variables := map[string]interface{}{
// 		"repositoriesCursor": (*githubv4.String)(nil),
// 		"languagesCursor":    (*githubv4.String)(nil),
// 	}
// 	var q struct {
// 		User struct {
// 			ID           githubv4.ID
// 			Repositories struct {
// 				Nodes    []Repository
// 				PageInfo struct {
// 					EndCursor   githubv4.String
// 					HasNextPage bool
// 				}
// 			} `graphql:"repositories(first: 50, after: $repositoriesCursor)"`
// 		} `graphql:"user(login: \"BTiberiu99\")"`
// 	}

// 	// Get comments from all pages.
// 	var allRepositories []Repository
// 	for {
// 		err := client.Query(context.Background(), &q, variables)
// 		if err != nil {
// 			return err
// 		}
// 		allRepositories = append(allRepositories, q.User.Repositories.Nodes...)
// 		if !q.User.Repositories.PageInfo.HasNextPage {
// 			break
// 		}
// 		variables["repositoriesCursor"] = githubv4.NewString(q.User.Repositories.PageInfo.EndCursor)
// 	}
// 	reps := make([]RepositoryFinal, len(allRepositories))
// 	for i := range allRepositories {
// 		reps[i].Name = allRepositories[i].Name
// 		reps[i].DatabaseID = allRepositories[i].DatabaseID
// 		reps[i].URL = allRepositories[i].URL
// 		reps[i].IsPrivate = allRepositories[i].IsPrivate
// 		reps[i].PrimaryLanguage = allRepositories[i].PrimaryLanguage
// 		reps[i].License = allRepositories[i].LicenseInfo
// 		// //Useful data
// 		// PrimaryLanguage Language
// 		for {
// 			err := client.Query(context.Background(), &q, variables)
// 			if err != nil {
// 				return err
// 			}
// 			reps[i].Languages = append(reps[i].Languages, allRepositories[i].Languages.Nodes...)
// 			if !q.User.Repositories.PageInfo.HasNextPage {
// 				break
// 			}
// 			variables["languagesCursor"] = githubv4.NewString(allRepositories[i].Languages.PageInfo.EndCursor)
// 		}
// 	}

// 	err = client.Query(context.Background(), &q, variables)
// 	if err != nil {
// 		// Handle error.
// 		fmt.Println(reps)
// 	} else {
// 		fmt.Println(allRepositories)
// 	}

// 	// fmt.Println(q)
// 	return nil
// }

// func repositoriLanguages(user models.UserPlatform, reps []Repository) error {
// 	// variables := map[string]interface{}{
// 	// 	"languagesCursor": (*githubv4.String)(nil),
// 	// }

// 	// for i := range reps {
// 	// 	for {
// 	// 		err := client.Query(context.Background(), &q, variables)
// 	// 		if err != nil {
// 	// 			return err
// 	// 		}
// 	// 		allRepositories = append(allRepositories, q.User.Repositories.Nodes...)
// 	// 		if !q.User.Repositories.PageInfo.HasNextPage {
// 	// 			break
// 	// 		}
// 	// 		variables["repositoriesCursor"] = githubv4.NewString(q.User.Repositories.PageInfo.EndCursor)
// 	// 	}
// 	// }
// 	return nil
// }

// func projectLicense(PlatformID string, ID int64) error {
// 	return nil
// }
