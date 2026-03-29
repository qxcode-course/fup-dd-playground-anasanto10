package main
import "fmt"
func main() {
    var hora, min, dia, mes, ano int16
    fmt.Scan(&hora, &min, &dia, &mes, &ano)
    ano = ano % 100

    fmt.Printf("%02d:%02d %02d/%02d/%02d\n", hora, min, dia, mes, ano)
}
