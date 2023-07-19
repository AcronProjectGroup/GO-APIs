/*

The approach Go takes for strings may seem odd, but it has its uses. Bytes are important when you
care about storing strings and you need to know how much space to allocate. Characters are important
when you are concerned with the contents of strings, such as when inserting a new character into an
existing string.
Both facets of strings are important. It is important, however, to understand whether you need to deal
with bytes or characters for any given operation.
You might be tempted to work only with bytes, which will work as long as you use only those characters
that are represented by a single byte, which typically means ASCII. This may work at first, but it almost
always ends badly, specifically when your code processes characters entered by a user with a non-
ASCII character set or processes a file containing non-ASCII data. For the small amount of extra work
required, it is simpler and safer to accept that Unicode does really exist and rely on Go to deal with
translating bytes into characters.

*/
package main