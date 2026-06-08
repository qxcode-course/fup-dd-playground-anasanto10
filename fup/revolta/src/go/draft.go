package main
import "fmt"

func criarEpreencher() []int {
    qtd := 0
    fmt.Scan(&qtd)
    arr := make([]int, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    return arr
}

func main() {
    arr := criarEpreencher()
    par := 0
    impar := 0

    for i := 0; i < len(arr); i++ {
        if arr[i]%2 == 0 {
            par += arr[i]
        }
        if arr[i]%2 != 0 {
            impar += arr[i]
        }
    }
    if impar > par {
        fmt.Println("soldados")
    } else if par > impar {
        fmt.Println("rebeldes")
    } else {
        fmt.Println("empate")
    }
}