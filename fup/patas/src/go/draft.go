package main
import "fmt"

func criarEpreencher() []string {
    qtd := 0
    fmt.Scan(&qtd)
    arr := make([]string, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    return arr
}

func main() {
    var cb, cc int
    fmt.Scan(&cb, &cc)
    arr := criarEpreencher()

    total := 0
    for i := 0; i < len(arr); i++{
        if arr[i] == "v"{
            total += 4
        } else if arr[i] == "g"{
            total += 2
        } else if arr[i] == "c"{
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
