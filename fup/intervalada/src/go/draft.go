package main
import "fmt"
func main() {
    var n, li, ls int
    fmt.Scan(&n, &li, &ls)

    arr := make([]int, n)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    intervalo := 0
    
    for i := 0; i < n; i++ {
        if arr[i] >= li && arr[i] <= ls {
            intervalo++
        }
    }
    fmt.Println(intervalo)
}
