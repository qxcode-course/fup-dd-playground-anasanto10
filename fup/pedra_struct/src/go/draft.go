package main
import "fmt"

type Jogada struct {
    p1 int
    p2 int
}

func calcPont (jogadas Jogada) (bool, int) {
    if jogadas.p1 < 10 || jogadas.p2 < 10 {
        return false, 0 
    } 
    pont := jogadas.p1 - jogadas.p2

    if pont < 0 {
        pont = -pont
    }
    return true, pont
}

func melhorJogada (jogadas []Jogada) int {
    for i := 1; i < len(jogadas); i++ {
        value, menor := calcPont(jogadas[0])
        if jogadas[i] < menor {
            menor = jogadas[i] 
        }
    }
}
func main() {
    var qtd int
    fmt.Scan(&qtd)

    jogadas := make([]Jogada, qtd)
    for i := range jogadas {
        fmt.Scan(&jogadas[i].p1, &jogadas[i].p2)
    }
}