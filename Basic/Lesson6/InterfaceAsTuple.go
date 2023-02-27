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

	var i InterfaceHere = otherStruct{3, 3}
	//fmt.Println(i.Sum())
	fmt.Printf("(%v,%T)\n",i,i ) // интерфейс представляет из себя значение (3,3) и тип main.otherStruct как тип интерфейса - () добавлены как для кортежа из питона

}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}

