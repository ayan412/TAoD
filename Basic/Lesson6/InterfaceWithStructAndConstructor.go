package main

import "fmt"

var _ User = &superUser{}

type superUser struct {
	Name      string
	Age       int
	isBlocked bool
}

func (s *superUser) Block() {
	s.isBlocked = true
}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
	isBlocked           bool
}

func (u *user) Block() {
	u.isBlocked = true
}

type User interface {
	Block()
}

func NewUser(fio, address, phone string) User {
	u := user{} // или user{}
	return &u
}

func main() {
	u := NewUser("Kim", "33", "9379992")
	u.Block()
	fmt.Println(u)

}
