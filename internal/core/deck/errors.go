package deck

import "errors"

var (
	ErrCreateDeck   = errors.New("unable to create deck")
	ErrDeckNotFound = errors.New("unable to find deck")
	ErrUpdateDeck   = errors.New("unable to update deck")
)

// SvcError is a custom error type that holds both internal and application errors.
type SvcError struct {
	InternalErr error
	AppErr      error
}

func NewSvcError(internalErr, appErr error) SvcError {
	return SvcError{
		InternalErr: internalErr,
		AppErr:      appErr,
	}
}

// Error is implementation of error interface.
func (x SvcError) Error() string {
	if x.InternalErr == nil && x.AppErr == nil {
		return ""
	}
	return errors.Join(x.InternalErr, x.AppErr).Error()
}
