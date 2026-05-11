package main
import "fmt"
func main() {
    var cb, cc, qtd int
    fmt.Scan(&cb, &cc, &qtd)

    animais := make([]string, qtd)

    for i := 0; i < qtd; i++{
        fmt.Scan(&animais[i])
    }

    total := 0

    for i := 0; i < qtd; i++{
        if animais[i] == "v"{
            total += 4
        } else if animais[i] == "g"{
            total += 2
        } else if animais[i] == "c"{
            total += 4
        } 
    }
    fmt.Println(total)

    dif_chico := cb - total
    if dif_chico < 0{
        dif_chico = -dif_chico
    }
    dif_ceb := cc - total
    if dif_ceb < 0{
        dif_ceb = -dif_ceb
    }

    if dif_chico < dif_ceb {
        fmt.Println("Chico Bento")
    } else if dif_ceb < dif_chico {
        fmt.Println("Cebolinha")
    } else {
        fmt.Println("empate")
    }
}
