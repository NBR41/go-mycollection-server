package model

import (
	"database/sql"
)

func InsertUser(nickname, email, password string, db *sql.DB) (*int64, error) {
	res, err := db.Exec(
		`
INSERT INTO users (id, nickname, email, password, ts_create, ts_update)
VALUES (null, ?, ?, ?, NOW(), NOW())
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
	return &id, err
}

func UpdateUserNickname(id int64, nickname string, db *sql.DB) error {
	_, err := db.Exec(
		`UPDATE users set nickname = ?, ts_update = NOW() where id = ?`,
		nickname, id,
	)
	return err
}

func UpdateUserPassword(id int64, password string, db *sql.DB) error {
	_, err := db.Exec(
		`UPDATE users set password = ?, ts_update = NOW() where id = ?`,
		password, id,
	)
	return err
}

func DeleteUser(id int64, db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM users where id = ?`, id)
	return err
}
