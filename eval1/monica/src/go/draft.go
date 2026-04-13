package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)

    c := m - (a + b)
    
    if a > b && a > c {
        fmt.Println(a)
    } else if b > a && b > c {
        fmt.Println(b)
    } else {
        fmt.Println(c)
    }
}
