package main 

var _ User = &superUser{}

type superUser struct {
	Name string
	Age int

}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

func (u *user) ChangeFIO(newFio string) {
	u.FIO = newFio
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
}

func NewUser(fio, address, phone string) User {
	u := user {
		FIO: fio,
		Address: address,
		Phone: phone,
	}

	return &u
}

func main() {


}

 