package main
import "fmt"
func main() {
    var r int
    fmt.Scan(&r)

    valor := make([]float64, r)
    for i := range valor {
        fmt.Scan(&valor[i])
    }

    chutes := make([]float64, r)
    for i := range chutes {
        fmt.Scan(&chutes[i])
    }

    escolhas := make([]string, r)
    for i := range escolhas {
        fmt.Scan(&escolhas[i])
    }

    jog1 := 0
    jog2 := 0 

    for i := 0; i < r; i++{
        if chutes[i] == valor[i] ||
         valor[i] > chutes[i] && escolhas[i] == "m" ||
         valor[i] < chutes[i] && escolhas[i] == "M" {
            jog1++
        } else {
            jog2++
        }
    }

    if jog1 > jog2 {
        fmt.Println("primeiro")
    } else if jog2 > jog1 {
        fmt.Println("segundo")
    } else {
        fmt.Println("empate")
    }
}