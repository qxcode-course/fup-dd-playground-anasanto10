package main
import "fmt"
func main() {
    var chute1, valor_real float64
    var escolha string
    fmt.Scan(&chute1, &escolha, &valor_real)

    if chute1 == valor_real {
        fmt.Println("primeiro")
    } else if (escolha == "M" && chute1 > valor_real) ||
              (escolha == "m" && chute1 < valor_real) {
                fmt.Println("primeiro")
              } else {
                fmt.Println("segundo")
              }
}
