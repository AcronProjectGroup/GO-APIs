dlv debug main.go

break bp1 main.main:3
# Result
#   Breakpoint 1 set at 0x697716 for main.main() c:/tools/main.go:8

condition bp1 i == 2

continue
# Result
# Hello, Go
# 0
# 1
# > [bp1] main.main() ./01-Preparing for Debugging.go:8 (hits goroutine(1):1 total:1) (PC: 0x49cc45)
#      3: import "fmt"
#      4:
#      5: func main() {
#      6:         fmt.Println("Hello, Go")
#      7:         for i := 0; i < 5; i++ {
# =>   8:                 fmt.Println(i)
#      9:         }
#     10: }
#     11:
#     12:
#     13:

