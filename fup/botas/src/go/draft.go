package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    var esquerda[61]int
    var direita[61]int

    for i := 0; i < N; i++{
        var tamanho int
        var lado string
        fmt.Scan(&tamanho, &lado)

        if lado == "E" {
            esquerda[tamanho]++
        } else {
            direita[tamanho]++
        }
    }

    pares := 0

    for i := 0; i < 61; i++ {
        if esquerda[i] < direita[i]{
            pares += esquerda[i]
        } else {
            pares += direita[i]
        }
    }
    fmt.Println(pares)
}