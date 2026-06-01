package main
import "fmt"
func main() {
    var n, m int
    fmt.Scan(&n, &m)

    inicio := 1
    for i := 1; i < n; i++{
        inicio *= 10
    }

    fim := inicio*10 -1

    var lista []int

    for i := inicio; i <= fim; i++{
        for j := i; j <= fim; j++{
            produto := i * j

            original := produto
            invertido := 0
            aux := produto

            for aux > 0{
                invertido = invertido*10 + aux%10
                aux /= 10
            }

            if original == invertido {
                lista = append(lista, produto)
            }
        }
    }

    for i := 0; i < len(lista)-1; i++{
        for j := 0; j < len(lista)-1; j++{
            if lista[j] < lista[j+1]{
                lista[j], lista[j+1] = lista[j+1], lista[j]
            }
        }
    }

    for i:=0; i < m && i < len(lista); i++{
        fmt.Println(lista[i])
    }
}