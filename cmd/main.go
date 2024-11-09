package main

import (
	"fmt"
	"validate"
)

type User struct {
	Name  string `validate:"required,min=2,max=32"`
	Email string `validate:"required,email"`
	Age   int    `validate:"min=18"`
}

func main() {
	user := User{
		Name:  "John",
		Email: "john@gmail.com",
		Age:   17,
	}

	ok, errors := validate.Validate(user)
	if !ok {
		fmt.Println(errors)
	}
}
