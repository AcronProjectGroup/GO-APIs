/*

The Go linter will report an error for any feature
that is exported from a package and that has not been
described in a comment.

Comments should be simple and descriptive, and the convention is to begin
the comment with the name of the feature, like this:




*/

// Product describes an item for sale

package main

type Product struct {
	Name, Category string // Name and type of the product
	price          float64
}
