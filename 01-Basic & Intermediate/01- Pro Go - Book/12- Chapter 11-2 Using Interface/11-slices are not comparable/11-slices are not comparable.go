/*

Adding a Field in the service.go File in the methodsAndServices Folder


package main
type Service struct {
	description string
	durationMonths int
	monthlyFee float64
  

	features []string

}

slices are not comparable. Compile and execute the project, and you will see

the effect of the new field:

panic: runtime error: comparing uncomparable type main.Service
goroutine 1 [running]:
main.main()
        C:/main.go:20 +0x1c5
exit status 2


*/

package main