package db

import (
	"database/sql"
	"fmt"
)

// Reset database
func Reset(driverName, dataSource, dbname string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	if err = resetDB(db, dbname); err != nil {
		return err
	}
	return db.Close()
}

// Migrate database
func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	if err = createPhoneNumsTable(db); err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, dbname string) error {
	statement := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return createDB(db, dbname)
}

func createDB(db *sql.DB, dbname string) error {
	statement := fmt.Sprintf("CREATE DATABASE %s", dbname)
	_, err := db.Exec(statement)
	return err
}

func createPhoneNumsTable(db *sql.DB) error {
	statement := `CREATE TABLE IF NOT EXISTS phone_numbers(
		id SERIAL,
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
}

func (db *DB) insertPhone(phoneNum string) (int, error) {
	statement := "INSERT INTO phone_numbers (value) VALUES ($1) RETURNING id"
	var id int
	err := db.db.QueryRow(statement, phoneNum).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
