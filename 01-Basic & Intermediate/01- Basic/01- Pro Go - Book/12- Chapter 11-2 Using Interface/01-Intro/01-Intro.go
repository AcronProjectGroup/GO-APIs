/*


the Go type rules prevent them from being used together, such as
creating a slice that contains both types of value.


This problem is addressed using interfaces
which describe a set of methods without specifying the
implementation of those methods.

If a type implements all the methods defined by the interface, then a
value of that type can be used wherever the interface is permitted.

Interfaces describe only methods and not fields.

*/


package main
