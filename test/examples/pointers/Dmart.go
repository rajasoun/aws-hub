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
