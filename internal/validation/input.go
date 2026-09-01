package validation

import (
	"errors"
	"os"
)

func ValidateInput(args []string) error {
	if len(os.Args) < 2 {
		return errors.New("provide filename")
	}

	return nil
}
