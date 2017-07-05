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
	rows, err := m.db.Query(`SELECT user_id, nickname, email, verified, admin FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []User{}
	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.IsVerified, &u.IsAdmin); err != nil {
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
		&u.ID, &u.Nickname, &u.Email, &u.IsVerified, &u.IsAdmin,
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
	return m.getUser(
		`SELECT user_id, nickname, email, activated, admin FROM users WHERE user_id = ?`,
		id,
	)
}

// GetUserByEmailOrNickname returns user by email or nickname
func (m *Model) GetUserByEmailOrNickname(email, nickname string) (*User, error) {
	return m.getUser(
		`SELECT user_id, nickname, email, activated, admin FROM users WHERE email = ? OR nickname = ?`,
		email, nickname,
	)
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (m *Model) GetAuthenticatedUser(password, email, nickname string) (*User, error) {
	return m.getUser(
		`
SELECT user_id, nickname, email, activated, admin
FROM users
WHERE password = ? AND (email = ? OR nickname =?)`,
		password, email, nickname,
	)
}

// InsertUser insert user
func (m *Model) InsertUser(nickname, email, password string) (*User, error) {
	res, err := m.db.Exec(
		`
INSERT INTO users (user_id, nickname, email, password, activated, admin, create_ts, update_ts)
VALUES (null, ?, ?, ?, 0, 0, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
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
		`UPDATE users set nickname = ?, update_ts = NOW() where user_id = ?`,
		nickname, id,
	)
	return err
}

// UpdateUserPassword updates user password by ID
func (m *Model) UpdateUserPassword(id int64, password string) error {
	_, err := m.db.Exec(
		`UPDATE users set password = ?, update_ts = NOW() where user_id = ?`,
		password, id,
	)
	return err
}

// UpdateUserActivation update user activation by ID
func (m *Model) UpdateUserActivation(id int64, activated bool) error {
	_, err := m.db.Exec(
		`UPDATE users set activated = ?, update_ts = NOW() where user_id = ?`,
		activated, id,
	)
	return err
}

// DeleteUser deletes user by ID
func (m *Model) DeleteUser(id int64) error {
	_, err := m.db.Exec(`DELETE FROM users where user_id = ?`, id)
	return err
}

// InsertBook inserts book
func (m *Model) InsertBook(name string) (*Book, error) {
	res, err := m.db.Exec(
		`
INSERT INTO books (book_id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
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
	return m.getBook(`SELECT book_id, name from books where id = ?`, id)
}

// GetBookList returns book list
func (m *Model) GetBookList() ([]Book, error) {
	rows, err := m.db.Query(`SELECT book_id, name FROM books`)
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
		`UPDATE books set name = ?, update_ts = NOW() where book_id = ?`,
		name, id,
	)
	return err
}

// DeleteBook delete book by ID
func (m *Model) DeleteBook(id int64) error {
	_, err := m.db.Exec(`DELETE FROM books where book_id = ?`, id)
	return err
}

// GetOwnershipList returns book list by user ID
func (m *Model) GetOwnershipList(userID int64) ([]Ownership, error) {
	rows, err := m.db.Query(
		`SELECT b.book_id, b.name FROM ownerships u JOIN books b USING(book_id) where user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []Ownership{}
	for rows.Next() {
		b := Book{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		b.initURL()
		ub := Ownership{
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

// GetOwnership returns user book association
func (m *Model) GetOwnership(userID, bookID int64) (*Ownership, error) {
	b, err := m.getBook(
		`
SELECT b.book_id, b.name
FROM ownerships u
JOIN books b USING(book_id)
where u.user_id = ? and b.book_id = ?`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, nil
	}
	ub := &Ownership{UserID: userID, BookID: bookID, Book: b}
	ub.initURL()
	return ub, nil
}

// InsertOwnership inserts user book association
func (m *Model) InsertOwnership(userID, bookID int64) (*Ownership, error) {
	_, err := m.db.Exec(
		`
INSERT INTO ownerships (user_id, book_id, create_ts, update_ts)
VALUES (?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	ub := &Ownership{UserID: userID, BookID: bookID}
	ub.initURL()
	return ub, nil
}

func (m *Model) UpdateOwnership(userID, bookID int64) error {
	_, err := m.db.Exec(
		`UPDATE ownerships set update_ts = NOW() where user_id = ? and book_id = ?`,
		userID, bookID,
	)
	return err
}

// DeleteOwnership deletes user book association
func (m *Model) DeleteOwnership(userID, bookID int64) error {
	_, err := m.db.Exec(`DELETE FROM ownerships where user_id = ? and book_id = ?`, userID, bookID)
	return err
}
