package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    pecas := ((n+1) * (n+2))/2

    fmt.Println(pecas)
}
