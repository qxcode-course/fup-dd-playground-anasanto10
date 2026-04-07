package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    if a == b{
        fmt.Println("Empate")
    } else {
        diff := (a - b + 15) % 15

        if diff >= 1 && diff <= 7 {
            fmt.Println("Jogador 2")
        } else {
            fmt.Println("Jogador 1")
        }
    }
}
