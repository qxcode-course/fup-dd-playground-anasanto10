package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    cartas := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&cartas[i])
    }

    fmt.Print("[")

    for i := 0; i < N; i++ {
        if cartas[i] == 1 {
            fmt.Print("A")
        } else if cartas[i] == 11 {
            fmt.Print("J")
        } else if cartas[i] == 12 {
            fmt.Print("Q")
        } else if cartas[i] == 13 {
            fmt.Print("K")
        } else {
            fmt.Print(cartas[i])
        }

        if i != N-1 {
            fmt.Print(", ")
        }
    }

    fmt.Println("]")
}
