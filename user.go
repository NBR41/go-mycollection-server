package main

import (
	"fmt"
)

type user struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	URL         string `json:"url,omitempty"`
	IsValidated bool
	IsAdmin     bool
}

func (u *user) initURL() {
	u.URL = fmt.Sprintf("/users/%d", u.ID)
}
