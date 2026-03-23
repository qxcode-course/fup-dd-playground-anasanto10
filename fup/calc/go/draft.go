package main
import "fmt"
func main() {
    var a, b int16
    var op string
    fmt.Scan(&a, &b, &op)
    if (op == "+") {
        fmt.Println(a + b)
    } else if (op == "-") {
        fmt.Println(a - b)
    } else if (op == "*") {
        fmt.Println(a * b)
    } else if (op == "/"){
        fmt.Println(a / b)
    }
    
}
