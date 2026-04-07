package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    fim := b
    fmt.Print("[ ")
    for {
        fmt.Print(a, b, " ")
        a++
        b--
        if a > fim {
            break
        }
    }
    fmt.Println("]")
}
