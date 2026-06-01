package main
import "fmt"
func main() {
    arr := make([]int, 5)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    maior := arr[0]
    menor := arr[0]

    for i := 0; i < len(arr); i++{

        if arr[i] > maior {
            maior = arr[i]
        }

        if arr[i] < menor {
            menor = arr[i]
        }
    }

    soma := maior + menor
    fmt.Println(soma)

}