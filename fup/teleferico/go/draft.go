package main
import "fmt"
func main() {
    var C, A int
    fmt.Scan(&C, &A)

    capacidade := A / (C - 1)

    if A % (C - 1) != 0 {
        capacidade = capacidade + 1
        fmt.Println(capacidade)
    } else {
        capacidade = A / (C - 1)
        fmt.Println(capacidade)
    }
}
