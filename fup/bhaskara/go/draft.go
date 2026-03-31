package main
import "fmt"
import "math"
func main() {
    var A, B, C float64
    fmt.Scan(&A, &B, &C)

    delta := (B*B) - 4*A*C

    if delta < 0 {
        fmt.Println("nao ha raiz real")
    } else if delta == 0 {
        x := -B / (2 *A)
        if x == 0 {
            x = 0 
        }
        fmt.Printf("%.2f\n", x)
    } else {
        raizDelta := math.Sqrt(delta)
        x1 := (-B + raizDelta) / (2 * A)
        x2 := (-B - raizDelta) / (2 * A)
        fmt.Printf("%.2f\n", x1)
        fmt.Printf("%.2f\n", x2)
    }
}
