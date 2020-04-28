package actions

import (
	"management/mailers"

	"github.com/gobuffalo/buffalo/worker"
	"github.com/pkg/errors"
)

var (
	noTypeEmail          = errors.New("No type email provided")
	typeEmailNonExistent = errors.New("Task for type email dosen't exist")
)

func SendEmail(args worker.Args) error {

	var (
		typeEmail string
		sendEmail func(worker.Args) error
	)

	if val, ok := args["type_email"].(string); !ok {
		return noTypeEmail
	} else {
		typeEmail = val
	}

	if val, ok := mailers.Mails[typeEmail]; !ok {
		return typeEmailNonExistent
	} else {
		sendEmail = val
	}

	return sendEmail(args)

}

func init() {
	w := App().Worker // Get a ref to the previously defined Worker
	w.Register("send_email", SendEmail)
}
