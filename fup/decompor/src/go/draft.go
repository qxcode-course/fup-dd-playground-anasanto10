package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    vet := make([]int, 0)

    for n > 0 {
        resto := n%10
        vet = append(vet, resto)
        n = n/10
    }

    i := 0
    j := len(vet)-1

    for i < j {
        vet[i], vet[j] = vet[j], vet[i]
        i++
        j--
    }
    for i := 0; i < len(vet); i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(vet[i])
    }
    fmt.Println()
}