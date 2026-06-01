package main
import "fmt"
func main() {
    var dist, n int
    fmt.Scan(&dist, &n)

    arr := make([]int, n)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    
}
