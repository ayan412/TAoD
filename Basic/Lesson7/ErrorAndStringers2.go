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

func main() {
	err := method1()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")

}

func method1() error {
	return method2()
}

func method2() error {
	return method3() 
}

func method3() error {
// business logic implementaion
	return nil //fmt.Errorf("some error")
}