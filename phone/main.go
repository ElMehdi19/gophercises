package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	phonedb "github.com/ElMehdi19/gophercises/phone/db"
	_ "github.com/lib/pq"
)

var (
	host, _     = os.LookupEnv("DBHOST") //"127.0.0.1"
	user, _     = os.LookupEnv("DBUSER")
	password, _ = os.LookupEnv("DBPASS")
	dbname, _   = os.LookupEnv("DBNAME")
)

func main() {
	must(ensureEnvs(host, user, password, dbname))
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", host, user, password)
	must(phonedb.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Seed())

	phones, err := db.FindAllPhones()
	must(err)

	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		normalized := normalize(p.Value)
		if p.Value != normalized {
			existing, err := db.FindPhone(normalized)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p))
			} else {
				p.Value = normalized
				must(db.UpdatePhone(p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

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

func ensureEnvs(envs ...string) error {
	for _, env := range envs {
		if len(env) == 0 {
			return errors.New("must set environement variables: DBHOST, DBUSER, DBPASS, DBNAME")
		}
	}
	return nil
}
