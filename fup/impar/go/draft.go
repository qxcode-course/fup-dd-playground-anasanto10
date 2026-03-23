package main
import "fmt"
func main() {
    var p, D1, D2 int
    fmt.Scan(&p, &D1, &D2)
    total := D1 + D2
    if (total%2 == 0){
        fmt.Println(p)
    } else {
        fmt.Println(1 - p)
    }
    
}
