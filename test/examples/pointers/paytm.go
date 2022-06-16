package pointers

import (
	"errors"
)

var ErroEmptyList = errors.New("empty list")

//we have user struct
type User struct {
	Name string
	list []string
}

//creates new user
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

//add's items in the list
func (u *User) AddItem(item string) {
	if len(item) > 0 {
		u.list = append(u.list, item)
	}
}

//return my list, if list is empty returns error
func (u *User) GiveMyList() ([]string, error) {
	if len(u.list) == 0 {
		return nil, ErroEmptyList
	}
	return u.list, nil
}

// //func main() {

// 	//RohiniList := NewUser("Rohini")
// 	SomePersonList := NewUser("Unknown")

// 	RohiniList.AddItem("Chocolates")
// 	RohiniList.AddItem("Mobile")

// 	SomePersonList.AddItem("Guns")
// 	SomePersonList.AddItem("walkie-talkie")

// 	log.Println("Rohini's items:")
// 	RohiniList.ShowList()

// 	abcList := *RohiniList
// 	abcList.Name = "abc"

// 	xyzList := RohiniList
// 	xyzList.Name = "xyz"

//}

//x = "hi" => memeory address 0X02537182
// x = "hello" = > memory address 0X62735283
// y =&x
// y = hello

// x := &a >& = address
// x > pointer variable
// y = *x > y

//&
//*
