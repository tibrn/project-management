package mailers

import (
	"fmt"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

func SendActivateAccounts(to string, token string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Activate Account"
	m.From = "barontiberiu@gmail.com"
	m.To = []string{to}
	err := m.AddBody(r.HTML("activate_account.html"), render.Data{
		"activate_path": envy.Get("APP_URL", "") + "/activate/email?token=" + token,
	})
	if err != nil {
		fmt.Println(err)
		return errors.WithStack(err)
	}

	fmt.Println(m)
	fmt.Println(smtp)
	return smtp.Send(m)
}
