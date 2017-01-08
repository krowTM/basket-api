package basketapi

import "testing"

func TestValidateRegister(t *testing.T) {
	user := User{Email: "dsdsd"}
	valid, err := user.ValidateRegister()
	if err == nil {
		t.Error("ValidateRegister email failed. err should not be nil")
	}
	if valid {
		t.Error("ValidateRegister email failed. valid should false")
	}

	user = User{Email: "dsdsd@tete.ro"}
	valid, err = user.ValidateRegister()
	if err == nil {
		t.Error("ValidateRegister failed. err should not be nil")
	}
	if valid {
		t.Error("ValidateRegister failed. valid should false")
	}

	user = User{Email: "dsdsd@tete.ro", FirstName: "f name"}
	valid, err = user.ValidateRegister()
	if err == nil {
		t.Error("ValidateRegister failed. err should not be nil")
	}
	if valid {
		t.Error("ValidateRegister failed. valid should false")
	}

	user = User{Email: "dsdsd@tete.ro", FirstName: "f name", LastName: "l name", Password: "password"}
	valid, err = user.ValidateRegister()
	if err != nil {
		t.Error("ValidateRegister failed. err should be nil")
	}
	if !valid {
		t.Error("ValidateRegister failed. valid should true")
	}
}

func TestValidateLogin(t *testing.T) {
	user := User{Email: "dsdsd"}
	valid, err := user.ValidateLogin()
	if err == nil {
		t.Error("ValidateLogin email failed. err should not be nil")
	}
	if valid {
		t.Error("ValidateLogin email failed. valid should false")
	}

	user = User{Email: "dsdsd@tete.ro"}
	valid, err = user.ValidateLogin()
	if err == nil {
		t.Error("ValidateLogin failed. err should not be nil")
	}
	if valid {
		t.Error("ValidateLogin failed. valid should false")
	}

	user = User{Email: "dsdsd@tete.ro", Password: "Password"}
	valid, err = user.ValidateLogin()
	if err != nil {
		t.Error("ValidateLogin failed. err should be nil")
	}
	if !valid {
		t.Error("ValidateLogin failed. valid should true")
	}
}
