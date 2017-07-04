package main

import (
	"fmt"
)

// User struct for users
type User struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	URL         string `json:"url,omitempty"`
	IsValidated bool
	IsAdmin     bool
}

func (u *User) initURL() {
	u.URL = fmt.Sprintf("/users/%d", u.ID)
}

// Book struct for books
type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

func (b *Book) initURL() {
	b.URL = fmt.Sprintf("/books/%d", b.ID)
}

// UserBook struct for user book association
type UserBook struct {
	UserID int64  `json:"user_id"`
	BookID int64  `json:"book_id"`
	URL    string `json:"url,omitempty"`
	Book   *Book  `json:"book,omitempty"`
}

func (u *UserBook) initURL() {
	u.URL = fmt.Sprintf("/users/%d/books/%d", u.UserID, u.BookID)
}
