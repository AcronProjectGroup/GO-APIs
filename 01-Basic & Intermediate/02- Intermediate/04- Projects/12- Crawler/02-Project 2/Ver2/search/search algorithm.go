package search

import (
	"github.com/boltdb/bolt"
	"strings"
)

type Searcher struct {
	db *bolt.DB
}

func NewSearcher(db *bolt.DB) *Searcher {
	return &Searcher{
		db: db,
	}
}

func (s *Searcher) Search(query string) []string {
	// Process the query
	words := NewQueryProcessor().Process(query)
	var results []string
	
	// Open a read-only transaction
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("IndexBucket"))
	
		// For each word, get the corresponding URLs from the index
		for _, word := range words {
			val := b.Get([]byte(word))
			if val != nil {
				urls := strings.Split(string(val), ",")
				results = append(results, urls...)
			}
		}
		return nil
	})
	
	// Return the results
	return results
}
