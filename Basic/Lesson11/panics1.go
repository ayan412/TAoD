package main

import (
	"errors"
	"fmt"
	"log"
)

type AppError struct {
	Message string
	Err error
}

func (ae *AppError) Error() string {
	return ae.Message
}

func main() {
	divide(4, 0)
	fmt.Println("after panic")

}

func divide(a, b int) {
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("this is AppError handle panic", err)
				} else {
					fmt.Println("this is Other Error handle panic ")
			}
			default:
				panic("some panic")
		}
			log.Println("panic happened:", err)
		}
	}()
	fmt.Println(div(a, b))
}

func div(a, b int) int {
	if b == 0 {
		panic(&AppError{
			Message: "this is divide by zero custom error!!!",
			Err: nil,
		})
	}
	return a / b
}