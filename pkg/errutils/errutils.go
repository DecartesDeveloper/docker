package errutils

import "github.com/pkg/errors"

// WrapNil does the same as errors.Wrap but wraps even if e is nil
func WrapNil(e, with error) error {
	if e == nil {
		return with
	}
	if with == nil {
		return e
	}
	return errors.Wrap(e, with.Error())
}
