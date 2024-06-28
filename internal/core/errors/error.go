package errors

import "fmt"

func Repository(msg string, err error) error {
	// [repository 'database'] function not exist
	return fmt.Errorf("[repository '%s'] %v ", msg, err)
}

func Service(msg string, err error) error {
	// [service 'product'] [repository 'database'] function not exist
	return fmt.Errorf("[service '%s'] %v ", msg, err)
}
