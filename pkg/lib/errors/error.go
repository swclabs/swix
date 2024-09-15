// Package errors are returned error messages layer
package errors

import (
	"fmt"
)

// Repository function return error message
func Repository(msg string, err error) error {
	if err == nil {
		return nil
	}
	// [repos 'database'] function not exist
	return fmt.Errorf("[repos '%s'] %v ", msg, err)
}

// Service function return error message
func Service(msg string, err error) error {
	if err == nil {
		return nil
	}
	// [service 'product'] [repos 'database'] function not exist
	return fmt.Errorf("[service '%s'] %v ", msg, err)
}
