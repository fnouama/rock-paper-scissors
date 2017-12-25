package model

import "testing"

func TestNewPlayer(t *testing.T) {

	p := NewPlayer("David", "Cooper")

	if p.UserName != "davidcooper" {
		t.Error("The username should be `davidcooper`")
	}

	if p.FirstName != "David" {
		t.Error("The first name should be `David`")
	}

	if p.LastName != "Cooper" {
		t.Error("The last name should be `Cooper`")
	}

	if p.Email != "david.cooper@gmail.com" {
		t.Error("The email should be `david.cooper@gmail.com`")
	}
}
