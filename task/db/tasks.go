package db

import (
	"encoding/binary"
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

//returns the key of created task or an error
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := intToByte(int(id64))

		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, err
}

//intToByte returns an 8-byte big endian representation of v
func intToByte(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b

}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
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
