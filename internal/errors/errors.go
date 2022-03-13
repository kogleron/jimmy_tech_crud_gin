package errors

import "errors"

func New(text string, code int) CodeError {
	return &codeTextError{
		err:  errors.New(text), //nolint:goerr113
		code: code,
	}
}

func WithCode(err error, code int) CodeError {
	return &codeTextError{
		err:  err,
		code: code,
	}
}

type CodeError interface {
	Error() string
	Code() int
}

type codeTextError struct {
	err  error
	code int
}

func (e *codeTextError) Code() int {
	return e.code
}

func (e *codeTextError) Error() string {
	return e.err.Error()
}
