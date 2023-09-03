/*

Function			Description
---------			-----------------------------------------------------------------
Replace				This function alters the string s by replacing occurrences of the string old with the
(s, old, new, n)	string new. The maximum number of occurrences that will be replaced is specified by
					the int argument n.

ReplaceAll			This function alters the string s by replacing all occurrences of the string old with
(s, old, new)		the string new. Unlike the Replace function, there is no limit on the number of
					occurrences that will be replaced.

Map(func, s)		This function generates a string by invoking the custom function for each character in
					the string s and concatenating the results. If the function produces a negative value,
					the current character is dropped without a replacement.

*/

package intro