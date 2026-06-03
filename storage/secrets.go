/*
Secrets.go is where the operations with the DB resides.
*/
package storage

import (
	"database/sql"
	"fmt"
)

// Secret struct, they are equal to what the rows of the DB save.
type Secret struct {
	ID int
	Name string 
	Key string
	Key_Length int
}

// Insert method accepts a name, a key, and a hashed Master Key, returns a int64 that is
// the id of the new row and can return a error aswell, the method is to register a new
// secret within the DB.
func (db *DB) Insert(name, key string, mk []byte) (int64, error) {
	// Encrypts the passed key.
	encrypted_key, err := Encrypt(key, mk)
	if err != nil {
		return 0, err
	}

	// Exec a SQL command to save the new secret.
	res, err := db.conn.Exec(
		`INSERT INTO secrets (name, key, key_length) VALUES (?, ?, ?)`,
		name, encrypted_key, len(key),
		)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// List method returns a slice of instantiate Secrets, it goes trough every secret on DB
// saving then.
func (db *DB) List() ([]Secret, error) {
	// Query the db.
	rows, err := db.conn.Query(`SELECT id, name, key, key_length FROM secrets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var secrets []Secret
	// Iterate over them saving on the slice.
	for rows.Next() {
		var s Secret
		rows.Scan(&s.ID, &s.Name, &s.Key, &s.Key_Length)
		secrets = append(secrets, s)
	}

	return secrets, rows.Err()
}

// GetByName method accepts a string and then query it against the DB, if the name is
// found the method returns a instantiated Secret of it.
func (db *DB) GetByName(name string) (*Secret, error) {
	var s Secret
	err := db.conn.QueryRow(
		`SELECT id, name, key, key_length FROM secrets WHERE name = ?`,
		name,
		).Scan(&s.ID, &s.Name, &s.Key, &s.Key_Length)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Secret %q not found.", name)
	}
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Delete method deletes a row based on the id of that is passed.
func (db *DB) Delete(id int) error {
	_, err := db.conn.Exec(`DELETE FROM secrets WHERE id = ?`, id)
	return err
}

// Update function method accepts a id of the row that will be changed, a Secret Struct 
// with the new content and the master key in order to encrypt key if needed, return error
// if any.
func (db *DB) Update(id_to_change int, new_secret Secret, mk []byte) error {
	encrypted_key, err := Encrypt(new_secret.Key, mk)
	if err != nil {
		return err
	}

	_, err = db.conn.Exec(`UPDATE secrets SET name = ? key = ? key_length = ? WHERE id = ?`,
		new_secret.Name, encrypted_key, len(new_secret.Key), id_to_change)

	return err
}

// AddMasterKey hash the argument passed and insert it into the 'config' table, returns it's 
// id if no error.
func (db *DB) AddMasterKey(key string) (int64, error) {
	hashed_key := HashMasterKey(key)
	res, err := db.conn.Exec(
		`INSERT INTO config (name, key) VALUES (?, ?)`,
		"master_key", hashed_key,
		)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetHashedMasterKey method get's the master key.
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

// MasterKeyExists verify if Master Key exists.
func (db *DB) MasterKeyExists() (bool, error) {
	var mk Secret
	err := db.conn.QueryRow(
		`SELECT id, name, key FROM config WHERE name = ?`,
		"master_key",
		).Scan(&mk.ID, &mk.Name, &mk.Key)

	if err == sql.ErrNoRows {
		return false, sql.ErrNoRows
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

/*
Index:
type Secret struct
func (db *DB) Insert(name, key string, mk []byte) (int64, error)
func (db *DB) List() ([]Secret, error)
func (db *DB) GetByName(name string) (*Secret, error) 
func (db *DB) Delete(id int) error
func (db *DB) Update(id_to_change int, new_secret Secret, mk []byte) error
func (db *DB) AddMasterKey(key string) (int64, error)
func (db *DB) GetHashedMasterKey() (*Secret, error)
func (db *DB) MasterKeyExists() (bool, error)
*/
