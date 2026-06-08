package main
import "fmt"
func main() {
    var a1, op, a2 string
    fmt.Scan(&a1, &op, &a2)

    x := int(a1[0] - 'a')
    y := int(a2[0] - 'a')

    var resultado int
    if op == "+" {
        resultado = (x + y) % 26
    } else {
        resultado = (x - y + 26) % 26
    }

    fmt.Println(string(byte(resultado + 'a')))
}