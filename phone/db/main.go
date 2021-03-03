package db

import (
	"database/sql"
	"fmt"
)

// DB type
type DB struct {
	db *sql.DB
}

// Phone relation schema
type Phone struct {
	ID    int
	Value string
}

// Open database connection
func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Close database connection
func (db *DB) Close() error {
	return db.db.Close()
}

// Seed add entries to phone_numbers table
func (db *DB) Seed() error {
	data := []string{
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	for _, p := range data {
		if _, err := db.insertPhone(p); err != nil {
			return err
		}
	}
	return nil
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

// FindPhone fetch one phone based on value column
func (db *DB) FindPhone(phoneNumber string) (*Phone, error) {
	statement := "SELECT * FROM phone_numbers WHERE value=$1"
	var phone Phone
	err := db.db.QueryRow(statement, phoneNumber).Scan(&phone.ID, &phone.Value)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &phone, err
}

// FindAllPhones fetch all rows from phone_numbers
func (db *DB) FindAllPhones() ([]Phone, error) {
	statement := "SELECT * FROM phone_numbers"
	rows, err := db.db.Query(statement)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var phones []Phone
	for rows.Next() {
		var phone Phone
		err := rows.Scan(&phone.ID, &phone.Value)
		if err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return phones, nil
}

// DeletePhone delete phone from phone_numbers
func (db *DB) DeletePhone(phone Phone) error {
	statement := "DELETE FROM phone_numbers WHERE id=$1"
	_, err := db.db.Exec(statement, phone.ID)
	return err
}

// UpdatePhone update value column
func (db *DB) UpdatePhone(p Phone) error {
	statement := "UPDATE phone_numbers SET value=$2 WHERE id=$1"
	_, err := db.db.Exec(statement, p.ID, p.Value)
	return err
}

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

func createPhoneNumsTable(db *sql.DB) error {
	statement := `CREATE TABLE IF NOT EXISTS phone_numbers(
		id SERIAL,
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
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
