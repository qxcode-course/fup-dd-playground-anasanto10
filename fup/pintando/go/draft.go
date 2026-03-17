package main
import "fmt"
import "math"
func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    var per = (a + b + c) / 2 
    var area = math.Sqrt(per * (per - a) * (per - b) * (per - c))
    fmt.Printf("%.2f\n", area )
}
