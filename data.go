package flashcards

import (
	"os"

	"github.com/coreos/bbolt"
)

type data struct {
	db     *bolt.DB
	dbName string
}

func newData() *data {
	return &data{
		db:     new(bolt.DB),
		dbName: "flashcards.db",
	}
}

func (d *data) setValue(bucket, key, value string) {
	d.db, _ = bolt.Open(d.dbName, os.ModePerm, nil)
	defer d.db.Close()

	d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			b, _ = tx.CreateBucketIfNotExists([]byte(bucket))
		}
		return b.Put([]byte(key), []byte(value))
	})
}
