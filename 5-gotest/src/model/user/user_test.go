package user

import (
	"testing"
)

func TestUser(t *testing.T) {

	nome := "Jhon"

	u := NewUser()

	if u.GetName() != "" {
		t.Errorf("user.Name failed, expected %v, got %v", "", u.GetName())
	}

	u.SetName(nome)

	if u.GetName() != nome {
		t.Errorf("user.Name failed, expected %v, got %v", nome, u.GetName())
	}
}
