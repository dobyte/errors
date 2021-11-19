package main

import (
    "encoding/json"
    "fmt"
    "os"
    "unsafe"
    
    "github.com/dobyte/errors"
    "github.com/dobyte/errors/example/errcode"
)

type response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Stack   string `json:"stack,omitempty"`
}

func main() {
    if err := controller(); err != nil {
        failed(err)
    }
    
    success()
}

func failed(err error) {
    var e = new(customError)
    errors.As(err, &e)
    
    writeJsonExit(response{
        Code:    e.code,
        Message: e.message,
        Stack:   fmt.Sprintf("%+v", err),
    })
}

func success() {
    writeJsonExit(response{
        Code:    errcode.Success.Int(),
        Message: errcode.Success.String(),
    })
}

func writeJsonExit(res response) {
    b, _ := json.Marshal(res)
    fmt.Println(*(*string)(unsafe.Pointer(&b)))
    os.Exit(0)
}

func controller() error {
    if err := service(); err != nil {
        return err
    }
    
    return nil
}

func service() error {
    if err := dao(); err != nil {
        return errors.WithError(err, NewError(errcode.BusinessError))
    }
    
    return nil
}

func dao() error {
    return errors.WrapError(NewError(errcode.InternalServerError))
}