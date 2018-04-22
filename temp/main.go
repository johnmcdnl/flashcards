package main

import (
	"fmt"
	"os"

	"github.com/coreos/bbolt"
	"github.com/sirupsen/logrus"
)

func main() {
	data := newData()
	bucket := "myBucket"
	data.create(bucket, "myKey", "myValue")
	data.readAll(bucket)
}

type data struct {
	db *bolt.DB
}

func newData() *data {
	d := new(data)
	d.db = new(bolt.DB)
	return d
}

func (d *data) open() {
	//d.close()
	db, _ := bolt.Open("flashcards.db", os.ModePerm, nil)
	d.db = db
}

func (d *data) close() {
	d.db.Close()
}

func (d *data) create(bucket, key, value string) error {
	return d.update(bucket, key, value)
}

func (d *data) read() {
	defer d.close()
	d.open()

}

func (d *data) readAll(bucket string) {
	defer d.close()
	d.open()

	tx := func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	}

	d.db.View(tx)

}

func (d *data) update(bucket, key, value string) error {
	defer d.close()
	d.open()

	tx := func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			logrus.Debugln("(d *data) update() CreateBucketIfNotExists %s", bucket)
			b, _ = tx.CreateBucketIfNotExists([]byte(bucket))
		}

		return b.Put([]byte(key), []byte(value))
	}

	return d.db.Update(tx)
}

func (d *data) delete() {}
