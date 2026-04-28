package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)

    arr := make([]int, qtd)
    for i := 0; i < qtd; i++ {
        fmt.Scan(&arr[i])
    }
    for i := 0; i < qtd; i++ {
        fmt.Println(arr[i])
}
    if qtd == 0{
        fmt.Println()
 }
}