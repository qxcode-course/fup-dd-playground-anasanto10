package main
import "fmt"
func main() {
    var nome string
    fmt.Scan(&nome)
    var soma int32 = 0
    for _, c := range nome {
        soma += c
    }
    fmt.Println(soma % 50)
}
