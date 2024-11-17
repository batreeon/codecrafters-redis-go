package constant

import "errors"

var (
	ErrInvaildInput     error = errors.New("invalid input")
	ErrParameterMissing error = errors.New("parameter missing")
)
