package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go say("hello", ch) // запуск метода say в отдельном потоке через ключ-е слово go
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println(<-ch)

}

func say(word string, ch chan string) {
	time.Sleep(10 * time.Second)
	fmt.Println(word)
	ch <- "exit"
}
