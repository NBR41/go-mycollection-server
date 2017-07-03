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

func (m *model) Close() error {
	return m.db.Close()
}

func (m *model) GetUserByEmail(email string) (*user, error) {
	var u = user{}
	err := m.db.QueryRow(
		`SELECT id, nickname, email, activated FROM users WHERE email = ?`,
		email,
	).Scan(&u.ID, u.Nickname, u.Email, u.IsValidated)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *model) GetAuthenticatedUser(password, email, nickname string) (*user, error) {
	var u = user{}
	err := m.db.QueryRow(
		`
SELECT id, nickname, email, activated
FROM users
WHERE password = ? AND (email = ? OR nickname =?)`,
		password, email, nickname,
	).Scan(&u.ID, u.Nickname, u.Email, u.IsValidated)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *model) InsertUser(nickname, email, password string) (*user, error) {
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

func (m *model) UpdateUserNickname(id int64, nickname string) error {
	_, err := m.db.Exec(
		`UPDATE users set nickname = ?, ts_update = NOW() where id = ?`,
		nickname, id,
	)
	return err
}

func (m *model) UpdateUserPassword(id int64, password string) error {
	_, err := m.db.Exec(
		`UPDATE users set password = ?, ts_update = NOW() where id = ?`,
		password, id,
	)
	return err
}

func (m *model) UpdateUserActivation(id int64, activated bool) error {
	_, err := m.db.Exec(
		`UPDATE users set activated = ?, ts_update = NOW() where id = ?`,
		activated, id,
	)
	return err
}

func (m *model) DeleteUser(id int64) error {
	_, err := m.db.Exec(`DELETE FROM users where id = ?`, id)
	return err
}
