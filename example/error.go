package main

type ErrCode interface {
    Int() int
    String() string
}

type customError struct {
    code    int
    message string
}

func (e *customError) Error() string { return e.message }

func NewError(code ErrCode) error {
    return &customError{code: code.Int(), message: code.String()}
}