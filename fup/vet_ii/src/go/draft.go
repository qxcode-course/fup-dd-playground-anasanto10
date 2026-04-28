package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)

    arr := make([]int, qtd)

    fmt.Printf("[")
    for i := 0; i < qtd; i++{
        fmt.Scan(&arr[i])        
    }
    for i := 0; i < qtd; i++{
        fmt.Print(" ", arr[i])
    }
    fmt.Println(" ]")
}
