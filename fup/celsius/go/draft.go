package main
import "fmt"
func main() {
    var t_celsius, T_f float64
    fmt.Scan(&t_celsius)
    T_f = (1.8 * t_celsius) + 32
    fmt.Printf("%.6f\n", T_f)
}
