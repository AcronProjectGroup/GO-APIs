package indexer

import (
	"github.com/boltdb/bolt"
)

type Indexer struct {
	// other fields
}

func NewIndexer( /* parameters */ ) *Indexer {
	return &Indexer{
		// initialization
	}
}

func (i *Indexer) Index(data Data) {
	// Process data
	// Store in BoltDB
}
