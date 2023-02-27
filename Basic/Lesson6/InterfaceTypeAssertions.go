package main 

import "fmt"

func main() {

	var a interface{} = 3.14
	// s := a.(string)
	// fmt.Println(s)

	// s, ok := a.(string)
	// fmt.Println(s, ok)

	// f, ok := a.(float64)
	// fmt.Println(f, ok)

	// // g := a.(float64) //panic
	// // fmt.Println(g) 

	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	case bool:
		fmt.Println("a is bool")
	default: 
		fmt.Printf("uknown type %T\n", a)
	}

 }