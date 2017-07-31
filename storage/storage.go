package storage

import (
	"encoding/binary"
	"log"

	"github.com/boltdb/bolt"
)

// Storage is generic storage, contains all game data access
type Storage interface {
	ProfileStorage
	HomeStorage
}

type boltStorage struct {
	db *bolt.DB
}

// New creates new boltDB storage
func New(filename string) Storage {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		log.Fatalf("can't open storage file %s: %q", filename, err)
	}
	return &boltStorage{
		db: db,
	}
}

func (bs *boltStorage) set(bucket []byte, id int64, data []byte) error {
	return bs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		return b.Put(bytesFromID(id), data)
	})
}

func (bs *boltStorage) get(bucket []byte, id int64) ([]byte, error) {
	data := make([]byte, 0)
	err := bs.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		data = b.Get(bytesFromID(id))
		return nil
	})
	return data, err
}

func bytesFromID(id int64) []byte {
	bytes := make([]byte, 8)
	binary.PutVarint(bytes, id)
	return bytes
}
