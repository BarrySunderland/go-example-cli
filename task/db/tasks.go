package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

//creates a bolt db file and creates a bucket if either doesn't already exist
func Setup(dbpath string) error {
	var err error
	db, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// func main() {

// 	// Open the my.db data file in your current directory.
// 	// It will be created if it doesn't exist.
// 	//insludes a timeout in case the file is locked by another process
// 	db, err := bolt.Open("task.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// }
