package main
import "fmt"
func main() {
    var r1, r2, r3, r4 string
    fmt.Scan(&r1, &r2, &r3, &r4)

    acertos := 0

    if r1 == "d" {
        acertos ++
    }
    if r2 == "a" {
        acertos ++
    }
    if r3 == "c" {
        acertos ++
    }
    if r4 == "d" {
        acertos ++
    }

    if acertos == 0 {
        fmt.Println("Nunca assistiu")
    } else if acertos == 1 {
        fmt.Println("Ja ouviu falar")
    } else if acertos == 2 {
        fmt.Println("Interessado no assunto")
    } else if acertos == 3 {
        fmt.Println("Fa")
    } else {
        fmt.Println("Super Fa")
    }
}
