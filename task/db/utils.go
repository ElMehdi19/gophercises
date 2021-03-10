package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

// CreateTask boltdb utility func to create new tasks
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

// ListTasks boltdb utility func to get all tasks from db
func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for key, val := c.First(); key != nil; key, val = c.Next() {
			tasks = append(tasks, Task{Key: btoi(key), Value: string(val)})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// DeleteTask boltdb utility func to delete a task
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(id int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
