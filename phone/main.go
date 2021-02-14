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
