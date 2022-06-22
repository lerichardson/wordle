package errors

import (
	"errors"
	"fmt"
)

var Red string = "\033[1;31m%s\033[0m"
var Black string = "\033[1;30m%s\033[0m"

func Handle(err error) {
	if err != nil {
		fmt.Printf("\n%sError: %e%s\n", Red, err, Black)
	}
}

func New(errorText string) (err error) {
	err = errors.New(errorText)
	return err
}
