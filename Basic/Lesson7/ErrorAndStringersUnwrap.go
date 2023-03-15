package main

import "fmt"

// redundant - встроен
// type error interface {
// 	Error() string
// }

type appError struct {
	Err    error // error встроенный интерфейс
	Custom string
	Field  int
}

func (e *appError) Error() string {
	return fmt.Sprintf("%s error", e.Custom)
}

func (e *appError) Unwrap() error {
	return e.Err
}

func main() {
	err := method1()
	if err != nil {
		fmt.Print(err.Unwrap())
		return
	}

	fmt.Println("success")

}

func method1() *appError {
	return method2()
}

func method2() *appError {
	return method3()
}

func method3() *appError {
	// business logic implementaion
	return &appError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong.",
	}
}
