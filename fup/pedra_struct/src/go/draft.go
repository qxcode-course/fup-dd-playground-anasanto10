package main
import "fmt"

type Jogada struct {
    p1 int
    p2 int
}

func calcPont (jogadas Jogada) (bool, int) {
    if jogadas.p1 < 10 || jogadas.p2 < 10 {
        return false, 1000
    } 
    pont := jogadas.p1 - jogadas.p2

    if pont < 0 {
        pont = -pont
    }
    return true, pont
}

func melhorJogada (jogadas []Jogada) int {
    anterior := 1000
    indice := 0
    for i := 0; i < len(jogadas); i++ {
        value, menor := calcPont(jogadas[i])
        if anterior > menor && value == true {
            anterior = menor
            indice = i
        }
    }
    return indice
}
func main() {
    var qtd int
    fmt.Scan(&qtd)

    jogadas := make([]Jogada, qtd)
    for i := range jogadas {
        fmt.Scan(&jogadas[i].p1, &jogadas[i].p2)
    }

    fmt.Println(melhorJogada(jogadas))
}