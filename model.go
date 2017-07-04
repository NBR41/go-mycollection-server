package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type model struct {
	db *sql.DB
}

func newModel(connString string) (*model, error) {
	db, err := sql.Open("mysql", connString+"?charset=utf8mb4,utf8")
	if err != nil {
		return nil, err
	}
	return &model{db: db}, nil
}

func (m *model) close() error {
	return m.db.Close()
}
func (m *model) getUserList() ([]user, error) {
	rows, err := m.db.Query(`SELECT id, nickname, email, activated, admin FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []user{}
	for rows.Next() {
		u := user{}
		if err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin); err != nil {
			return nil, err
		}
		l = append(l, u)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}
func (m *model) getUser(query string, params ...interface{}) (*user, error) {
	var u = user{}
	err := m.db.QueryRow(query, params...).Scan(
		&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin,
	)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *model) getUserByID(id int64) (*user, error) {
	return m.getUser(`SELECT id, nickname, email, activated, admin FROM users WHERE id = ?`, id)
}

func (m *model) getUserByEmailOrNickname(email, nickname string) (*user, error) {
	return m.getUser(
		`SELECT id, nickname, email, activated, admin FROM users WHERE email = ? OR nickname = ?`,
		email, nickname,
	)
}

func (m *model) getAuthenticatedUser(password, email, nickname string) (*user, error) {
	return m.getUser(
		`
	SELECT id, nickname, email, activated, admin
	FROM users
	WHERE password = ? AND (email = ? OR nickname =?)`,
		password, email, nickname,
	)
}

func (m *model) insertUser(nickname, email, password string) (*user, error) {
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
	return &user{ID: id, Email: email, Nickname: nickname}, err
}

func (m *model) updateUserNickname(id int64, nickname string) error {
	_, err := m.db.Exec(
		`UPDATE users set nickname = ?, ts_update = NOW() where id = ?`,
		nickname, id,
	)
	return err
}

func (m *model) updateUserPassword(id int64, password string) error {
	_, err := m.db.Exec(
		`UPDATE users set password = ?, ts_update = NOW() where id = ?`,
		password, id,
	)
	return err
}

func (m *model) updateUserActivation(id int64, activated bool) error {
	_, err := m.db.Exec(
		`UPDATE users set activated = ?, ts_update = NOW() where id = ?`,
		activated, id,
	)
	return err
}

func (m *model) deleteUser(id int64) error {
	_, err := m.db.Exec(`DELETE FROM users where id = ?`, id)
	return err
}

func (m *model) insertBook(name string) (*book, error) {
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
	return &book{ID: id, Name: name}, err
}

func (m *model) UpdateBook(id int64, name string) error {
	_, err := m.db.Exec(
		`UPDATE books set name = ?, ts_update = NOW() where id = ?`,
		name, id,
	)
	return err
}

func (m *model) deleteBook(id int64) error {
	_, err := m.db.Exec(`DELETE FROM books where id = ?`, id)
	return err
}

func (m *model) insertUserBook(userID, bookID int64) (*userBook, error) {
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
	return &userBook{UserID: userID, BookID: bookID}, err
}

func (m *model) updateUserBook(userID, bookID int64) error {
	_, err := m.db.Exec(
		`UPDATE users_books set ts_update = NOW() where user_id = ? and book_id = ?`,
		userID, bookID,
	)
	return err
}

func (m *model) deleteUserBook(userID, bookID int64) error {
	_, err := m.db.Exec(`DELETE FROM users_books where user_id = ? and book_id = ?`, userID, bookID)
	return err
}
