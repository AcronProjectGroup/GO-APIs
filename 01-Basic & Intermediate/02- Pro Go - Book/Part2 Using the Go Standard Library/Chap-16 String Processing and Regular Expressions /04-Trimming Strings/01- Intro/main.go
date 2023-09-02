/*

Function				Description
------------			------------------------------------------------------------------------------
TrimSpace(s)			This function returns the string s without leading or trailing whitespace
						characters.

Trim(s, set)			This function returns a string from which any leading or trailing characters
						contained in the string set are removed from the string s.

TrimLeft(s, set)		This function returns the string s without any leading character contained
						in the string set. This function matches any of the specified characters—use
						the TrimPrefix function to remove a complete substring.

TrimRight(s, set)		This function returns the string s without any trailing character contained
						in the string set. This function matches any of the specified characters—use
						the TrimSuffix function to remove a complete substring.

TrimPrefix(s, prefix)	This function returns the string s after removing the specified prefix string.
						This function removes the complete prefix string—use the TrimLeft
						function to remove characters from a set.

TrimSuffix(s, suffix)	This function returns the string s after removing the specified suffix string.
						This function removes the complete suffix string—use the TrimRight
						function to remove characters from a set.

TrimFunc(s, func)		This function returns the string s from which any leading or trailing
						character for which a custom function returns true are removed.

TrimLeftFunc(s, func)	This function returns the string s from which any leading character for
						which a custom function returns true are removed.

TrimRightFunc(s, func)	This function returns the string s from which any trailing character for
						which a custom function returns true are removed.


*/

package main