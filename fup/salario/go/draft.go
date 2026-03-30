package main
import "fmt"
func main() {
    var atual float32
    fmt.Scan(&atual)
    
    if atual <= 1000 {
        novo := (atual * 0.2) + atual
        fmt.Printf("%.2f\n", novo)
    } else if atual <= 1500 {
        novo := (atual * 0.15) + atual
        fmt.Printf("%.2f\n", novo)
    } else if atual <= 2000 {
        novo := (atual * 0.1) + atual
        fmt.Printf("%.2f\n", novo)
    } else {
        novo := (atual * 0.05) + atual
        fmt.Printf("%.2f\n", novo)
    }
}
