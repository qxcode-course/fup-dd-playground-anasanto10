package main
import "fmt"
func main() {
    var num1, num2 float32
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    fmt.Printf("%.1f\n", (num1+num2)/2)
}