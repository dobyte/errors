package errors

import (
	"fmt"
	"io"
)

type withStackMessage struct {
	msg string
	*stack
}

func (w *withStackMessage) Error() string { return w.msg }
func (w *withStackMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.msg)
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.msg)
	}
}

type withStackError struct {
	error
	*stack
}

func (w *withStackError) Cause() error               { return w.error }
func (w *withStackError) Unwrap() error              { return w.error }
func (w *withStackError) Is(target error) bool       { return Is(w.error, target) }
func (w *withStackError) As(target interface{}) bool { return As(w.error, target) }
func (w *withStackError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg }
func (w *withMessage) Cause() error  { return w.cause }
func (w *withMessage) Unwrap() error { return w.cause }
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", w.Cause())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, w.Error())
	}
}

type withError struct {
	cause error
	err   error
}

func (w *withError) Error() string              { return w.err.Error() }
func (w *withError) Cause() error               { return w.cause }
func (w *withError) Unwrap() error              { return w.cause }
func (w *withError) Is(target error) bool       { return Is(w.err, target) }
func (w *withError) As(target interface{}) bool { return As(w.err, target) }
func (w *withError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", w.Cause())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, w.Error())
	}
}

// WrapMessage Wrap a new error with message and stack.
func WrapMessage(message string) error {
	return &withStackMessage{message, callers()}
}

// WrapError Wrap a new error with error and stack.
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	return &withStackError{err, callers()}
}

// WithMessage Create a new error with message on top of the original error.
func WithMessage(cause error, message string) error {
	if cause == nil {
		return nil
	}
	switch cause.(type) {
	case *withStackMessage, *withStackError, *withMessage, *withError:
		return &withMessage{cause, message}
	default:
		return &withMessage{WrapError(cause), message}
	}
}

// WithError Create a new error with error on top of the original error.
func WithError(cause, err error) error {
	if cause == nil {
		return nil
	}
	switch cause.(type) {
	case *withStackMessage, *withStackError, *withMessage, *withError:
		return &withError{cause, err}
	default:
		return &withError{WrapError(cause), err}
	}
}

// Cause Recursively find the root cause of the error
func Cause(err error) error {
	type causer interface {
		Cause() error
	}
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
