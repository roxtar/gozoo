package gozoo

import "fmt"

type ZooError struct {
	errorCode ZookeeperError
}

func newZooError(code ZookeeperError) *ZooError {
	z := &ZooError{
		errorCode: code,
	}
	return z
}

func (z *ZooError) Error() string {
	if z != nil {
		return fmt.Sprintf("%s (%d)", z.errorCode, z.errorCode)
	}
	return ""
}
