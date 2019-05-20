package grifts

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/icrowley/fake"
	"github.com/markbates/grift/grift"
	"github.com/myWebsite/golang/actions"
	"github.com/myWebsite/golang/models"
	"github.com/pkg/errors"
)

//MAXINT64 ... MAX INT64 NUMBER
const MAXINT64 = 9223372036854775807

var users = []*models.User{}
var tasks = []*models.Task{}
var projects = []*models.Project{}
var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {

		for i := range languages {
			err := models.DB.Create(languages[i])
			fmt.Println(err)
			if err != nil {
				models.DB.Where("name = ?", languages[i].Name).First(languages[i])
			}
		}

		licens := []*models.License{}

		models.DB.All(&licens)

		for i := range licenses {
			err := models.DB.Create(licenses[i])
			if err == nil {
				licens = append(licens, licenses[i])
			}

		}

		licenses = licens

		for i := 0; i < 200; i++ {
			models.DB.Transaction(func(tx *pop.Connection) error {
				password := fake.Password(8, 15, true, true, true)
				user := &models.User{
					Email:                fake.EmailAddress(),
					PasswordPlain:        password,
					PasswordConfirmation: password,
					Name:                 fake.FirstName(),
					JoinedAt:             nulls.NewTime(time.Now()),
					Settings:             models.UserSetting{},
				}

				if _, err := user.Create(tx); err != nil {
					fmt.Println(err)
					return errors.WithStack(err)
				}

				verifyEmail := &models.UserVerify{
					UserID: user.ID,
					Type:   "activate-account",
					Token:  actions.GenerateToken(user.Slug.String + user.Email + fmt.Sprintf("%d", user.ID)),
				}

				_, err := tx.ValidateAndCreate(verifyEmail)

				if err != nil {
					fmt.Println(err)
					return err
				}

				users = append(users, user)

				nrAccounts := rand.Intn(6)
				for i := 0; i < nrAccounts; i++ {
					err := CreateUserPlatform(tx, user.ID)
					fmt.Println(err)
				}
				return nil
			})
		}

		return nil
	})

})

func CreateUserPlatform(tx *pop.Connection, UserID int64) error {
	platform := rand.Int63n(2) + 1
	username := fake.UserName()
	URL := ""
	if platform == 1 {
		URL += "https://api.github.com/users/"
	} else {
		URL += "https://api.gitlab.com/users/"
	}
	URL += username
	userPlatform := &models.UserPlatform{
		IDOnPlatform: rand.Int63n(MAXINT64),
		PlatformID:   platform,
		Username:     username,
		UserID:       UserID,
		URL:          URL,
		Token:        fake.CharactersN(40),
		TokenType:    "bearer",
		Limit:        5000,
	}

	_, err := tx.ValidateAndCreate(userPlatform)
	if err == nil {
		r := rand.Intn(7)
		for i := 0; i < r; i++ {
			CreateProject(tx, userPlatform)
		}
	}

	return nil
}

func CreateProject(tx *pop.Connection, usplatf *models.UserPlatform) error {

	project := &models.Project{
		PlatformID:   usplatf.PlatformID,
		IDOnPlatform: rand.Int63n(math.MaxUint32),
		Name:         fake.CharactersN(12),
		Description:  nulls.NewString(fake.CharactersN(250)),
		URL:          fake.CharactersN(50),
		LicenseID:    nulls.NewInt64(licenses[rand.Intn(len(licenses))].ID),
	}
	_, err := tx.ValidateAndCreate(project)
	if err == nil {
		nrTasks := rand.Intn(25)

		for i := 0; i < nrTasks; i++ {

			CreateTask(tx, project.ID)
		}

		nrLanguages := rand.Intn(5) + 2
		exist := map[string]int64{}
		for i := 0; i < nrLanguages; i++ {
			lang := languages[rand.Intn(len(languages))]
			_, ok := exist[lang.Name]
			for ok {
				lang = languages[rand.Intn(len(languages))]
				_, ok = exist[lang.Name]
			}

			exist[lang.Name] = lang.ID

			makeRelationProjectLanguage(tx, lang.ID, project.ID, i == 0)

		}

		exist2 := map[int64]bool{}

		nrUsers := rand.Intn(4)
		if len(users) < nrUsers {
			nrUsers = len(users)
		}
		for i := 0; i < nrUsers; i++ {
			user := users[rand.Intn(len(users))]
			_, ok := exist2[user.ID]
			for ok {
				user = users[rand.Intn(len(users))]
				_, ok = exist2[user.ID]
			}

			exist2[user.ID] = true

			makeRelationUserProject(tx, user.ID, project.ID)
		}

		projects = append(projects, project)
	}
	return nil
}

func CreateTask(tx *pop.Connection, projectID uuid.UUID) error {

	task := &models.Task{
		ProjectID:   projectID,
		Name:        fake.ProductName(),
		Description: fake.CharactersN(250),
		Progress:    float64(rand.Int63n(100)),
	}

	if _, err := tx.ValidateAndCreate(task); err != nil {
		fmt.Println(err)
		return err
	}

	exist := map[int64]bool{}

	usersTask := []*models.User{}
	nrUsers := rand.Intn(3)
	if len(users) < nrUsers {
		nrUsers = len(users)
	}
	if nrUsers > 0 {
		for i := 0; i < nrUsers; i++ {
			user := users[rand.Intn(len(users))]
			_, ok := exist[user.ID]
			for ok {
				user = users[rand.Intn(len(users))]
				_, ok = exist[user.ID]
			}

			exist[user.ID] = true
			usersTask = append(usersTask, user)

			makeRelationUserTask(tx, user.ID, task.ID)
		}

		nrComments := rand.Intn(10)
		for i := 0; i < nrComments; i++ {
			CreateComment(tx, usersTask[rand.Intn(len(usersTask))].ID, task.ID)
		}
	}

	tasks = append(tasks, task)
	return nil
}

func CreateComment(tx *pop.Connection, userID int64, taskID int64) error {

	comment := &models.Comment{
		UserID:  userID,
		TaskID:  taskID,
		Content: fake.CharactersN(250),
		Rating:  int8(rand.Intn(5)),
	}

	if _, err := tx.ValidateAndCreate(comment); err != nil {
		fmt.Println(err)

		return err
	}
	return nil
}

func makeRelationUserProject(tx *pop.Connection, userID int64, projectID uuid.UUID) error {
	relation := &models.UserProject{
		UserID:    userID,
		ProjectID: projectID,
	}

	if _, err := tx.ValidateAndCreate(relation); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}

func makeRelationProjectLanguage(tx *pop.Connection, languageID int64, projectID uuid.UUID, primary bool) error {
	relation := &models.ProjectLanguage{
		LanguageID: languageID,
		ProjectID:  projectID,
		Primary:    primary,
		Usage:      rand.Float64() * 99999999,
	}

	if _, err := tx.ValidateAndCreate(relation); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}

func makeRelationUserTask(tx *pop.Connection, userID int64, taskID int64) error {
	relation := &models.UserTask{
		TaskID: taskID,
		UserID: userID,
	}

	if _, err := tx.ValidateAndCreate(relation); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}
