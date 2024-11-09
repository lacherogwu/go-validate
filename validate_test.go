package validate_test

import (
	"testing"
	"validate"
)

type User struct {
	Name  string `validate:"required,min=2,max=32"`
	Email string `validate:"required,email"`
}

func TestValidateOk(t *testing.T) {

	u := User{
		Name:  "Serj",
		Email: "serj@gmail.com",
	}

	ok, errors := validate.Validate(u)
	if !ok {
		t.Errorf("Expected validation to pass, but got errors: %v", errors)
	}

}

func TestValidateFail(t *testing.T) {

	u := User{
		Name:  "S",
		Email: "john123@gmail.com",
	}

	ok, errors := validate.Validate(u)
	if ok {
		t.Errorf("Expected validation to fail, but got no errors: %v", errors)
	}
}
