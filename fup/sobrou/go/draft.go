package main
import "fmt"
func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    var preco1, preco2, preco3 float64
    fmt.Scan(&preco1, &preco2, &preco3)
    var saldo float64
    fmt.Scan(&saldo)
    var resto float64
    resto = (saldo - ((a*preco1) + (b*preco2) + (c*preco3)))
    fmt.Printf("%.2f\n", resto)
}
