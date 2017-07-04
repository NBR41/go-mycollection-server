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

type book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

func (b *book) initURL() {
	b.URL = fmt.Sprintf("/books/%d", b.ID)
}

type userBook struct {
	UserID int64  `json:"user_id"`
	BookID int64  `json:"book_id"`
	URL    string `json:"url,omitempty"`
}

func (u *userBook) initURL() {
	u.URL = fmt.Sprintf("/users/%d/books/%d", u.UserID, u.BookID)
}
