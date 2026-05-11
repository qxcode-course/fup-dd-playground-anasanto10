package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    sabor := make([]string, qtd)
    turno := make([]string, qtd)

    for i := 0; i < qtd; i++{
        fmt.Scan(&sabor[i])
        fmt.Scan(&turno[i])
    }

    choc := 0
    lim := 0
    man := 0
    tar := 0

    for i := 0; i < qtd; i++{
    if sabor[i] == "c" {
        choc ++
    } else if sabor[i] == "l"{
        lim ++
    }
    
    if turno[i] == "m" {
        man ++
    } else if turno[i] == "t"{
        tar ++
    }
    }

    if choc > lim {
        fmt.Println("c")
    } else if lim > choc {
        fmt.Println("l")
    } else {
        fmt.Println("empate")
    }

    if man < tar {
        fmt.Println("m")
    } else if tar < man{
        fmt.Println("t")
    } else {
        fmt.Println("empate")
    }
}
