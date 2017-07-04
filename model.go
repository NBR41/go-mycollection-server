package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Model struct for model
type Model struct {
	db *sql.DB
}

// NewModel returns new instance of model
func NewModel(connString string) (*Model, error) {
	db, err := sql.Open("mysql", connString+"?charset=utf8mb4,utf8")
	if err != nil {
		return nil, err
	}
	return &Model{db: db}, nil
}

func (m *Model) close() error {
	return m.db.Close()
}

// GetUserList returns user list
func (m *Model) GetUserList() ([]User, error) {
	rows, err := m.db.Query(`SELECT id, nickname, email, activated, admin FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []User{}
	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin); err != nil {
			return nil, err
		}
		u.initURL()
		l = append(l, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

func (m *Model) getUser(query string, params ...interface{}) (*User, error) {
	var u = User{}
	err := m.db.QueryRow(query, params...).Scan(
		&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		u.initURL()
		return &u, nil
	}
}

// GetUserByID returns user by ID
func (m *Model) GetUserByID(id int64) (*User, error) {
	return m.getUser(`SELECT id, nickname, email, activated, admin FROM users WHERE id = ?`, id)
}

// GetUserByEmailOrNickname returns user by email or nickname
func (m *Model) GetUserByEmailOrNickname(email, nickname string) (*User, error) {
	return m.getUser(
		`SELECT id, nickname, email, activated, admin FROM users WHERE email = ? OR nickname = ?`,
		email, nickname,
	)
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (m *Model) GetAuthenticatedUser(password, email, nickname string) (*User, error) {
	return m.getUser(
		`
	SELECT id, nickname, email, activated, admin
	FROM users
	WHERE password = ? AND (email = ? OR nickname =?)`,
		password, email, nickname,
	)
}

// InsertUser insert user
func (m *Model) InsertUser(nickname, email, password string) (*User, error) {
	res, err := m.db.Exec(
		`
INSERT INTO users (id, nickname, email, password, activated, ts_create, ts_update)
VALUES (null, ?, ?, ?, 0, NOW(), NOW())
ON DUPLICATE KEY UPDATE ts_update = VALUES(ts_update)`,
		nickname, email, password,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	u := &User{ID: id, Email: email, Nickname: nickname}
	u.initURL()
	return u, nil
}

// UpdateUserNickname updates user nickname by ID
func (m *Model) UpdateUserNickname(id int64, nickname string) error {
	_, err := m.db.Exec(
		`UPDATE users set nickname = ?, ts_update = NOW() where id = ?`,
		nickname, id,
	)
	return err
}

// UpdateUserPassword updates user password by ID
func (m *Model) UpdateUserPassword(id int64, password string) error {
	_, err := m.db.Exec(
		`UPDATE users set password = ?, ts_update = NOW() where id = ?`,
		password, id,
	)
	return err
}

// UpdateUserActivation update user activation by ID
func (m *Model) UpdateUserActivation(id int64, activated bool) error {
	_, err := m.db.Exec(
		`UPDATE users set activated = ?, ts_update = NOW() where id = ?`,
		activated, id,
	)
	return err
}

// DeleteUser deletes user by ID
func (m *Model) DeleteUser(id int64) error {
	_, err := m.db.Exec(`DELETE FROM users where id = ?`, id)
	return err
}

// InsertBook inserts book
func (m *Model) InsertBook(name string) (*Book, error) {
	res, err := m.db.Exec(
		`
INSERT INTO books (id, name, ts_create, ts_update)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE ts_update = VALUES(ts_update)`,
		name,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	b := &Book{ID: id, Name: name}
	b.initURL()
	return b, nil
}

func (m *Model) getBook(query string, params ...interface{}) (*Book, error) {
	var b = Book{}
	err := m.db.QueryRow(query, params...).Scan(&b.ID, &b.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		b.initURL()
		return &b, nil
	}
}

// GetBookByID returns book by ID
func (m *Model) GetBookByID(id int64) (*Book, error) {
	return m.getBook(`SELECT id, name from books where id = ?`, id)
}

// GetBookList returns book list
func (m *Model) GetBookList() ([]Book, error) {
	rows, err := m.db.Query(`SELECT id, name FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []Book{}
	for rows.Next() {
		b := Book{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		b.initURL()
		l = append(l, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// UpdateBook update book infos
func (m *Model) UpdateBook(id int64, name string) error {
	_, err := m.db.Exec(
		`UPDATE books set name = ?, ts_update = NOW() where id = ?`,
		name, id,
	)
	return err
}

// DeleteBook delete book by ID
func (m *Model) DeleteBook(id int64) error {
	_, err := m.db.Exec(`DELETE FROM books where id = ?`, id)
	return err
}

// GetUserBookList returns book list by user ID
func (m *Model) GetUserBookList(userID int64) ([]UserBook, error) {
	rows, err := m.db.Query(
		`SELECT b.id, b.name FROM users_books u JOIN books b ON (b.id = u.book_id) where user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []UserBook{}
	for rows.Next() {
		b := Book{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		b.initURL()
		ub := UserBook{
			UserID: userID,
			BookID: b.ID,
			Book:   &b,
		}
		ub.initURL()
		l = append(l, ub)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetUserBook returns user book association
func (m *Model) GetUserBook(userID, bookID int64) (*UserBook, error) {
	b, err := m.getBook(
		`SELECT b.id, b.name FROM users_books u JOIN books b ON (b.id = u.book_id) where user_id = ? and b.id = ?`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, nil
	}
	ub := &UserBook{UserID: userID, BookID: bookID, Book: b}
	ub.initURL()
	return ub, nil
}

// InsertUserBook inserts user book association
func (m *Model) InsertUserBook(userID, bookID int64) (*UserBook, error) {
	_, err := m.db.Exec(
		`
INSERT INTO users_books (user_id, book_id, ts_create, ts_update)
VALUES (?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE ts_update = VALUES(ts_update)`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	ub := &UserBook{UserID: userID, BookID: bookID}
	ub.initURL()
	return ub, nil
}

func (m *Model) UpdateUserBook(userID, bookID int64) error {
	_, err := m.db.Exec(
		`UPDATE users_books set ts_update = NOW() where user_id = ? and book_id = ?`,
		userID, bookID,
	)
	return err
}

// DeleteUserBook deletes user book association
func (m *Model) DeleteUserBook(userID, bookID int64) error {
	_, err := m.db.Exec(`DELETE FROM users_books where user_id = ? and book_id = ?`, userID, bookID)
	return err
}
