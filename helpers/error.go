package helpers

import "errors"

func BadRequestError(err error) error {
	return errors.New("bad request: " + err.Error())
}

func UnauthorizedError() error {
	return errors.New("unauthorized")
}
