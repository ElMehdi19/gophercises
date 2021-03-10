package db

import (
	"database/sql"
)

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
