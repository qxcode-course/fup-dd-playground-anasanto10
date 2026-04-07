package main
import "fmt"
func main() {
    var a, g, ra, rg float32
    fmt.Scan(&a, &g, &ra, &rg)

    custo_alcool := a / ra
    custo_gaso := g / rg

    if custo_alcool < custo_gaso {
        fmt.Println("A")
    } else {
        fmt.Println("G")
    }
}
