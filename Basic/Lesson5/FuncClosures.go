package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice // часть внешней ф-и totalPrice
	return func(x int) int { // anonymous function which generate closures
		sum += x
		return sum // sum возвращается в sum на 6-й строке или нет?
	}
}

func main() {
	orderPrice := totalPrice(1) // ч-з присваивание orderPrice стала ф-й к-я сама принимает и возвр-т число, поэтому при ее вызове ей перед-я число
	fmt.Println(orderPrice(1)) // поэтому при ее вызове ей передается число
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
}