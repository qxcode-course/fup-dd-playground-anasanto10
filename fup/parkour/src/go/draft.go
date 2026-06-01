package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    
}
