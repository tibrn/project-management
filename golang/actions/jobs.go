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

const (
	EmailJobArg = "type_email"
)

func SendEmail(args worker.Args) error {

	var (
		typeEmail string
		sendEmail func(worker.Args) error
	)
	//Check type_email args exist
	if val, ok := args[EmailJobArg].(string); !ok {
		return noTypeEmail
	} else {
		typeEmail = val
	}

	//Check type email job exist
	if val, ok := mailers.Mails[typeEmail]; !ok {
		return typeEmailNonExistent
	} else {
		sendEmail = val
	}

	//Send email
	return sendEmail(args)

}

func init() {
	w := App().Worker // Get a ref to the previously defined Worker
	w.Register("send_email", SendEmail)
}
