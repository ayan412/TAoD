package main

import "fmt"

// redundant - встроен
// type error interface {
// 	Error() string
// }

type AppError struct {
	Err    error // error встроенный интерфейс
	Custom string
	Field  int
}

func (e *AppError) Error() string {
	fmt.Println(e.Custom)
	return e.Err.Error()
}

func main () {
	err := m()
	if err != nil {
		fmt.Println(err.Error()) // реализован в методе Error() выше	
	}

}

func m() error {
	return &AppError{
		Err: fmt.Errorf("my error"),
		Custom: "value here",
		Field: 89,
	}
}