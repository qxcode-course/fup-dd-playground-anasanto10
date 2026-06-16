package main
import "fmt"

type restaurante struct {
        nomeRest string
        pontosRest int
    }

func escolherRest (restaurantes []restaurante) string {
    vencedor := restaurantes[0]

    for i := 1; i < len(restaurantes); i++ {
        if restaurantes[i].pontosRest > vencedor.pontosRest {
        vencedor = restaurantes[i]
     } else if restaurantes[i].pontosRest == vencedor.pontosRest {
        if restaurantes[i].nomeRest < vencedor.nomeRest {
            vencedor = restaurantes[i]
        }
      }
    }
    return vencedor.nomeRest
}
func main() {
    var n int
    fmt.Scan(&n)

    restaurantes := make([]restaurante, n)
    for i := range restaurantes {
        fmt.Scan(&restaurantes[i].nomeRest, &restaurantes[i].pontosRest)
    }

    fmt.Println(escolherRest(restaurantes))
}