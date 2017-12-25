package model

import (
	"time"
	"strings"
)

type Player struct {
	*User
	Created time.Time
}

func NewPlayer(firstName string, lastName string) *Player {
	return &Player{
		User: &User{
			UserName: strings.ToLower(firstName+lastName),
			FirstName:firstName,
			LastName:lastName,
			Email: strings.ToLower(firstName + "." + lastName) + "@gmail.com",
		},
		Created: time.Now(),

	}
}