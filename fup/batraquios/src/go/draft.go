package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
    vet1 := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&vet1[i])
    }

    var M int
    fmt.Scan(&M)
    vet2 := make([]int, M)

    for i := 0; i < M; i++ {
        fmt.Scan(&vet2[i])
    }


}
