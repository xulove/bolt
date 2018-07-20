package main

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
	"fmt"
)

func main() {
	// Open the database.
	db, err := bolt.Open("test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(db.Path())

	// Start a write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b, err := tx.CreateBucket([]byte("widgets"))
		if err != nil {
			return err
		}

		// Set the value "bar" for the key "foo".
		if err := b.Put([]byte("foo"), []byte("bar")); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	var valueByte []byte
	// Read value back in a different read-only transaction.
	if err := db.View(func(tx *bolt.Tx) error {
		valueByte = tx.Bucket([]byte("widgets")).Get([]byte("foo"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(valueByte))
	// Close database to release file lock.
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
