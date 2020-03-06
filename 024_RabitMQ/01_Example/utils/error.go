package utils

import (
	"github.com/pkg/errors"
	"github.com/rollbar/rollbar-go"
)

// FailOnError func
func FailOnError(err error, msg string) {
	if err != nil {
		rollbar.Critical(errors.WithMessage(err, msg))
		rollbar.Wait()
	}
}
