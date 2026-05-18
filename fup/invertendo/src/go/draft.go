package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    vetor := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&vetor[i])
    }

    for i := 0; i < N/2; i++ {
        aux := vetor[i]

        vetor[i] = vetor[N-1-i]

        vetor[N-1-i] = aux
    }
    
    fmt.Print("[ ")

    for i:= 0; i < N; i++ {
        fmt.Print(vetor[i])

        if i < N-1 {
            fmt.Print(" ")
        }
    }
    fmt.Print(" ]")
}
