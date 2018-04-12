package flashcards

import (
	"github.com/boltdb/bolt"
	"time"
	"os"
)

type data struct {
	db        *bolt.DB
	dbOptions *bolt.Options
	dbName    string
	buckets   []string
}

func newData() *data {
	data := &data{
		dbOptions: &bolt.Options{
			Timeout:         time.Duration(5 * time.Second),
			InitialMmapSize: 10000,
			ReadOnly:        false,
		},
		dbName:  "flashcards.db",
		buckets: []string{"commonNouns"},
	}

	for _, b := range data.buckets {
		data.createBucketIfNotExists(b)
	}

	return data
}

func (d *data) open() {
	db, _ := bolt.Open(d.dbName, os.ModePerm, d.dbOptions, )
	d.db = db
}
func (d *data) close() {
	d.db.Close()
}

func (d *data) create(bucket, key, value string) error {
	return d.update(bucket, key, value)
}

func (d *data) createBucketIfNotExists(bucket string) error {
	defer d.close()
	d.open()

	tx := func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
			return err
		}
		return nil
	}

	return d.db.Update(tx)
}

func (d *data) read(bucket, key string) string {

}

func (d *data) readAll(bucket string) []string {
	defer d.db.Close()
	d.open()

	var data []string
	tx := func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			data = append(data, string(v))
		}
		return nil
	}

	d.db.View(tx)

	return data
}

func (d *data) update(bucket, key, value string) error {
	defer d.close()
	d.open()

	tx := func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(value))
	}

	return d.db.Update(tx)
}

func (d *data) delete(bucket, key string) error {
	defer d.close()
	d.open()

	tx := func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	}

	return d.db.Update(tx)
}
