package search

import (
	"strings"
)

type QueryProcessor struct {
	// Your fields here
}


func NewQueryProcessor( /* parameters */ ) *QueryProcessor {
	return &QueryProcessor{
		// initialization
	}
}



func (qp *QueryProcessor) Process(query string) []string {
	// Convert the query to lowercase
	query = strings.ToLower(query)

	// Split the query into words
	words := strings.Fields(query)

	// Return the processed words
	return words
}
