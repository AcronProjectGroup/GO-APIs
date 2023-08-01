/*

...
func (product Product) printDetails() {
...
can be invoked like this:
...
Product.printDetails(Product{ "Kayak", "Watersports", 275 })
...




The name of the method’s receiver type, 
Product, in this case, is followed by a period and the method name. 
The argument is the Product value that will be used for the receiver value.
doesn’t apply when invoking a method through its receiver type, 
which means that a method with a pointer signature

*/

package main