/*
The Storage package is what handles all the DB operations, begining by creating and migrating
it to Insertions and Deletes.
*/
package storage

import (
	"database/sql"

	// Using SQLite on the moment, simple and easy to use, very light also, you don't
	// need a big and heavy SQL Database for just some passwords.
	_ "modernc.org/sqlite"
)

// DB struct, only field is the connection with the actual DB.
type DB struct {
	conn *sql.DB
}

// New function opens your DB connection and migrate if needed, it returns a initialized
// DB struct and a db method for migrating.
func New(path string) (*DB, error) {
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	db := &DB{conn}
	return db, db.migrate()
}

// migrate method create the DB tables if they don't exist, returns a error if any.
func (db *DB) migrate() error {
	_, err := db.conn.Exec(`
		CREATE TABLE IF NOT EXISTS secrets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			key TEXT NOT NULL,
			key_length INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS config (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			key TEXT NOT NULL
		)
	`)

	return err
}

// Close method is for closing the connection with the DB.
func (db *DB) Close() error {
	return db.conn.Close()
}
