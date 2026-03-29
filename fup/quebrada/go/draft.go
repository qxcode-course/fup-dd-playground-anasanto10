package main
import "fmt"
func main() {
    var n1, n2 int16
    fmt.Scan(&n1, &n2)
    div1 := n1 / n2
    resto := n1%n2
    div2 := float64(n1) / float64(n2)
    fmt.Println(div1)
    fmt.Println(resto)
    fmt.Printf("%.2f\n", div2)
}
