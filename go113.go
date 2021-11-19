package errors

import (
	stderrors "errors"
)

// New Wrapping for errors.New standard library
func New(msg string) error { return stderrors.New(msg) }

// Is Wrapping for errors.Is standard library
func Is(err, target error) bool { return stderrors.Is(err, target) }

// As Wrapping for errors.As standard library
func As(err error, target interface{}) bool { return stderrors.As(err, target) }

// Unwrap Wrapping for errors.Unwrap standard library
func Unwrap(err error) error { return stderrors.Unwrap(err) }
