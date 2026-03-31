package main
import "fmt"
func main() {
    var N, D, A int
    fmt.Scan(&N, &D, &A)

    botao := (D - A + N) % N
    fmt.Println(botao)
}
