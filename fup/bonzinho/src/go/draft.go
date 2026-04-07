package main
import "fmt"
func main() {
    var p1, p2, p3, trab float32
    fmt.Scan(&p1, &p2, &p3, &trab)
    
    menor := p1 
    if p2 < menor {
        menor = p2
    }
    if p3 < menor {
        menor = p3
    }

    soma_maiores := p1 + p2 + p3 - menor

    media := (trab + soma_maiores) / 3

    if media >= 7 {
        fmt.Printf("Aprovado com %.1f\n", media)
    } else {
        fmt.Printf("Final com %.1f\n", media)
    }
}
