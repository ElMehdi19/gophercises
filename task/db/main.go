package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task schema in boltdb
type Task struct {
	Key   string
	Value string
}

// Init initialize boltdb connection
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		// run transaction here
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
