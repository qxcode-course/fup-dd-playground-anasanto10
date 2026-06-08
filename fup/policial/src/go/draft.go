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

func ordernarCrescente(arr []int) []int {
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func main() {
    arr := criarEpreencher()
    ordernar := ordernarCrescente(arr)
    
    for i := 0; i < len(arr); i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(ordernar[i])
    }
    fmt.Println()
}