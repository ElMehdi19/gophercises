package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	user     = "postgres"
	password = "postgres"
	dbname   = "gophers_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", host, user, password)
	// db, err := sql.Open("postgres", psqlInfo)
	// must(err)
	// must(resetDB(db, dbname))
	// db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()
	must(createPhoneNumsTable(db))
	phones, err := getAllPhones(db)
	must(err)

	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		normalized := normalize(p.number)
		if p.number != normalized {
			existing, err := findPhone(db, normalized)
			must(err)
			if existing != nil {
				must(deletePhone(db, p))
			} else {
				p.number = normalized
				must(updatePhone(db, p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

}

type phone struct {
	id     int
	number string
}

func updatePhone(db *sql.DB, p phone) error {
	statement := "UPDATE phone_numbers SET value=$2 WHERE id=$1"
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func deletePhone(db *sql.DB, p phone) error {
	statement := "DELETE FROM phone_numbers WHERE id=$1"
	_, err := db.Exec(statement, p.id)
	return err
}

func findPhone(db *sql.DB, phoneNumber string) (*phone, error) {
	var p phone
	statement := "SELECT * FROM phone_numbers WHERE value=$1"
	err := db.QueryRow(statement, phoneNumber).Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func getAllPhones(db *sql.DB) ([]phone, error) {
	statement := "SELECT id, value FROM phone_numbers"
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	var phones []phone
	for rows.Next() {
		var id int
		var number string
		if err := rows.Scan(&id, &number); err != nil {
			return nil, err
		}
		phones = append(phones, phone{id, number})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return phones, nil
}

func getPhone(db *sql.DB, id int) (string, error) {
	statement := "SELECT value FROM phone_numbers WHERE id=$1"
	var phoneNumber string
	err := db.QueryRow(statement, id).Scan(&phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber, nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := "INSERT INTO phone_numbers (value) VALUES ($1) RETURNING id"
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
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

func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func normalize(phone string) string {
	var sb strings.Builder

	for _, char := range phone {
		if char >= '0' && char <= '9' {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}
