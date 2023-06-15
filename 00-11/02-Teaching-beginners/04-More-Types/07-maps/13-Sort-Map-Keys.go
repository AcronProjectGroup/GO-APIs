// Source : https://www.golangprograms.com/go-language/golang-maps.html



package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	unSortedMap := map[string]int{
		"India": 20,
		"Canada": 70, 
		"Germany": 15,
		"Sina": 29,
		"Mina": 30,
	}
 
	keys := make([]string, 0, len(unSortedMap))
 
	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
 
	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}

	fmt.Println(unSortedMap)
	fmt.Println(keys)

}





