// A slice doesn’t store any data, it just describes a section of an underlying array.
// یک برش هیچ داده ای را ذخیره نمی کند، فقط بخشی از یک آرایه زیرین را توصیف می کند.

// When you change an element of a slice, you modify the corresponding element of its underlying array, and other slices that share the same underlying array will see the change.
// هنگامی که یک عنصر از یک برش را تغییر می‌دهید، عنصر مربوط به آرایه زیرین آن را تغییر می‌دهید و سایر برش‌هایی که همان آرایه زیرین را به اشتراک می‌گذارند، این تغییر را مشاهده خواهند کرد.

/* Arraye:
		1- an ordered sequence of elements

*/


/* Slice:
		1- an ordered sequence of elements
		2- describes a section of an underlying array
		3- actualy any cheanges is on element of underlying array
		4-

*/



package main

import "fmt"

func main() {

	// Basic Create
		var s []int                   // a nil slice
		fmt.Println(s)
		// s[0] = 1
		
		s1 := []string{"foo", "bar"}
		fmt.Println(s1)

		s2 := make([]int, 2)          // same as []int{0, 0}
		fmt.Println(s2)
		
		
		s3 := make([]int, 2, 4)       // same as new([4]int)[:2]
		fmt.Println(s3)
		fmt.Println(len(s3), cap(s3)) // 2 4


	// create a slice by slicing an existing array or slice.
		a := [...]int{0, 1, 2, 3}     // an array
		s4 := a[1:3]               	  // s == []int{1, 2}        cap(s) == 3
		s4 = a[:2]                    // s == []int{0, 1}        cap(s) == 4
		s4 = a[2:]                    // s == []int{2, 3}        cap(s) == 2
		s4 = a[:]                     // s == []int{0, 1, 2, 3}  cap(s) == 4

		fmt.Println(s4)


}