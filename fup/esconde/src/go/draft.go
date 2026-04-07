package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    for i:= 0; i <= N; i++ {
        if i%2 != 0 {
            fmt.Println(i)
        }
    } 
    for i := N; i >= 0; i-- {
        if i%2 == 0 {
            fmt.Println(i)
        }
    } 
}
