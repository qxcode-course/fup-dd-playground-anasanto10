package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    risco := 0

    for i := 0; i < n; i++{
        
        if arr[i] == 0 {

            esquerda:= 0
            direita := 0

            if i > 0 {
                esquerda = arr[i-1]
            }

            if i < n-1 {
                direita = arr[i+1]
            }

            if esquerda != 1 && direita != 1 {
                risco++
            }
        }
    }
    fmt.Println(risco)
}
