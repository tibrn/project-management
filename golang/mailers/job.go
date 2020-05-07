package mailers

import "github.com/gobuffalo/buffalo/worker"

func NewJob(args worker.Args) worker.Job {
	return worker.Job{
		Queue:   "default",
		Handler: "send_email",
		Args:    args,
	}
}
