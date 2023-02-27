package main

import (
	"fmt"
)

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2

}

type InterfaceHere interface {
	Sum() int
}

func main() {

	var os *structHere
	var i InterfaceHere
	i = os	
	fmt.Printf("(%v,%T)\n",i,i ) // интерфейс представляет из себя значение (3,3) и тип main.otherStruct как тип интерфейса
	if i == nil { // постоянно nil т.к. structHere объявлена, но пуста и тип данных не nil,а *structHere
		fmt.Println("nil")
	}else{
		fmt.Println("not nill")
	}
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}

