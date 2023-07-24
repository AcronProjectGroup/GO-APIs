/*

Putting Structs in Context:




Question						Answer
-----------------------			---------------------------------------------------------
Are there any pitfalls or		Care must be taken to avoid unintentionally duplicating struct values and to
limitations?					ensure that fields that store pointers are initialized before they are used.

Are there any					will need to define custom types, for which structs are the only option.
alternatives?					Simple applications can use just the built-in data types, but most applications



Problem									Solution
----------								-----------------------------
Create a struct value					Use the literal syntax to create a new value and assign
										values to individual fields

Define a struct field whose type is		Define an embedded field
another struct

Compare struct values					Use the comparison operator, ensuring that the values
										being compared have the same type or types that have
										the same fields, all of which must be comparable

Convert struct types					Perform an explicit conversion, ensuring that the types
										have the same fields

Define a struct without assigning a		Define an anonymous struct
name

Prevent a struct from being duplicated	User a pointer
when it is assigned to a variable or
used as a function argument

Create struct values consistently		Define a constructor function





*/
package main