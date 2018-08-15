package db

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrDuplicateKey = errors.New("Duplicate Key")
)

const dbErrorUniqViolation = "23505"

// Takes the error from gorp and standardize it
func StandardizeError(err error) error {
	// Handle the PGSQL errors
	if e, ok := err.(*pq.Error); ok {
		switch e.Code {
		case dbErrorUniqViolation:
			return ErrDuplicateKey
		}
	}

	// Return the default error
	return err
}
