package storage

import (
	"database/sql"
	"fmt"
)

type Secret struct {
	ID int
	Name string 
	Key string
}

func (db *DB) Insert(name, key string) (int64, error) {
	res, err := db.conn.Exec(
		`INSERT INTO secrets (name, key) VALUES (?, ?)`,
		name, key,
		)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (db *DB) List() ([]Secret, error) {
	rows, err := db.conn.Query(`SELECT id, name, key FROM secrets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var secrets []Secret
	for rows.Next() {
		var s Secret
		rows.Scan(&s.ID, &s.Name, &s.Key)
		secrets = append(secrets, s)
	}

	return secrets, rows.Err()
}

func (db *DB) GetByName(name string) (*Secret, error) {
	var s Secret
	err := db.conn.QueryRow(
		`SELECT id, name, key FROM secrets WHERE name = ?`,
		name,
		).Scan(&s.ID, &s.Name, &s.Key)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Secret %q not found.", name)
	}
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (db *DB) Delete(id int) error {
	_, err := db.conn.Exec(`DELETE FROM secrets WHERE id = ?`, id)
	return err
}

func (db *DB) AddMasterKey(key string) (int64, error) {
	res, err := db.conn.Exec(
		`INSERT INTO secrets (name, key) VALUES (?, ?)`,
		"master_key", key,
		)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (db *DB) GetHashedMasterKey() (*Secret, error) {
	var mk Secret
	err := db.conn.QueryRow(
		`SELECT id, name, key FROM config WHERE name = ?`,
		"master_key",
		).Scan(&mk.ID, &mk.Name, &mk.Key)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Secret master key not found.")
	}
	if err != nil {
		return nil, err
	}

	return &mk, nil
}
