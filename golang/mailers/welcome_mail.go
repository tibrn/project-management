package mailers

import (
	"errors"
	"management/models"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/worker"
)

func SendWelcomeEmails(args worker.Args) error {

	var userID interface{}

	if val, ok := args["user_id"]; !ok {

		return errors.New("No user id provied")
	} else {
		userID = val
	}

	user := &models.User{}

	err := models.DB.Find(user, userID)

	if err != nil {
		return err
	}

	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Welcome Email"
	m.From = "barontiberiu@gmail.com"
	m.To = []string{user.Email}
	err = m.AddBody(r.HTML("welcome_email.html"), render.Data{
		"user": user,
	})

	if err != nil {
		return err
	}

	err = smtp.Send(m)

	return err
}
