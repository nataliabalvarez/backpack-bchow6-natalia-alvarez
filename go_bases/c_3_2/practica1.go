package main

import "fmt"

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Password string
}

func (user *User) ChangeName(newName string) {
	user.Name = newName
}
func (user *User) ChangeAge(newAge int) {
	user.Age = newAge
}
func (user *User) ChangeLastname(newLastname string) {
	user.Lastname = newLastname
}
func (user *User) ChangeEmail(newEmail string) {
	user.Email = newEmail
}
func (user *User) ChangePassword(newPassword string) {
	user.Password = newPassword
}

func main() {
	user := User{
		Name:     "Nat",
		Lastname: "Alvarez",
		Age:      31,
		Email:    "lala@mail.com",
		Password: "pass123",
	}

	fmt.Printf("User actual: %+v\n", user)

	user.ChangeName("Bele")
	user.ChangeEmail("Bele@lala.com")
	fmt.Printf("User modificado: %+v\n", user)

}
