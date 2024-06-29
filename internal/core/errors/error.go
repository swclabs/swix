package errors

import "fmt"

func Repository(msg string, err error) error {
	if err == nil {
		return nil
	}
	// [repository 'database'] function not exist
	return fmt.Errorf("[repository '%s'] %v ", msg, err)
}

func Service(msg string, err error) error {
	if err == nil {
		return nil
	}
	// [service 'product'] [repository 'database'] function not exist
	return fmt.Errorf("[service '%s'] %v ", msg, err)
}
