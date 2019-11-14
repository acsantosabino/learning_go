package user

import "fmt"

type User struct {
	name string
	age  int16
}

func NewUser() *User {
	return &User{name: "", age: 0}
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func Str(u *User) {
	fmt.Sprintln(&u.name)
}
