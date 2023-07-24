/*

Question				Answer
--------------			--------------------
What are they?			Methods are functions that are invoked on a struct and have access to all of the
						fields defined by the value’s type. Interfaces define sets of methods, which can be
						implemented by struct types.

Why are they useful?	These features allow types to be mixed and used through their common
						characteristics.

How are they used?		Methods are defined using the func keyword, but with the addition of a receiver.
						Interfaces are defined using the type and interface keywords.

Are there any pitfalls	Careful use of pointers is important when creating methods, and care must be taken
or limitations?			when using interfaces to avoid problems with the underlying dynamic types.

Are there any			These are optional features, but they make it possible to create complex data types
alternatives?			and use them through the common features they provide.



*/
package main