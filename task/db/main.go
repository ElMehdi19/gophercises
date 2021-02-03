package db

import (
	"time"

	"github.com/boltdb/bolt"
)

// Db boltdb instance
var Db *bolt.DB
var err error

func init() {
	Db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
}
