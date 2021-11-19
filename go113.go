package errors

import (
	stderrors "errors"
)

func New(msg string) error { return stderrors.New(msg) }

func Is(err, target error) bool { return stderrors.Is(err, target) }

func As(err error, target interface{}) bool { return stderrors.As(err, target) }

func Unwrap(err error) error { return stderrors.Unwrap(err) }
