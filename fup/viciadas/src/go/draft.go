package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)

    vet := make([]int, qtd)
    for i := range vet {
        fmt.Scan(&vet[i])
    }

    dif := 0 
    rep := make([]int, 0)
    for i := 0; i < len(vet); i++{
        if vet[i] != vet[i]+1 {
            dif++
        }
        if vet[i] == vet[i]+1 {
            rep = append(rep, vet[i])
        }
    }
    fmt.Println(dif, rep)
}