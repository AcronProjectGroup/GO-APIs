package indexer

import (
	"github.com/boltdb/bolt"
	"strings"
)

type Indexer struct {
	// other fields
	db *bolt.DB
}

func NewIndexer( /* parameters */ db *bolt.DB) *Indexer {
	return &Indexer{
		// initialization
		db: db,
	}
}

func (i *Indexer) Index(data Data) {
	// Process data
	// Store in BoltDB
	// Extract the text from the data
	// Split the text into words
	// For each word, add an entry in the index

	// ... Index the web pages ...
	// Open a writable transaction
	i.db.Update(func(tx *bolt.Tx) error {
		// For each word, add an entry in the database
		for word, urls := range index {
			// Use the word as the key and the URLs as the value
			err := tx.Bucket([]byte("IndexBucket")).Put([]byte(word),
				[]byte(strings.Join(urls, ",")))
			if err != nil {
				return err
			}
		}
		return nil
	})
}
