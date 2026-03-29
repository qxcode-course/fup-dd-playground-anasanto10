package main
import "fmt"
func main() {
    var vm, tempo, comb float32
    fmt.Scan(&vm, &tempo, &comb)
    tempo = tempo / 60
    dist := vm * tempo
    desempenho := dist / comb
    fmt.Printf("%.2f\n", desempenho)
}
