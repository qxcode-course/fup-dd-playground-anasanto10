package main
import "fmt"
func main() {
    var tam int
    fmt.Scan(&tam)

    arr := make([]int, tam)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    numero := 0 
    for _, valor := range arr {
        numero = numero*10 +valor
    }
    fmt.Println(numero)
}