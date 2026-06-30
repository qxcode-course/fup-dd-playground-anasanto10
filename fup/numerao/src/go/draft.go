package main
import "fmt"
func main() {
    var num string
    fmt.Scan(&num)

    somaPar := 0
    somaImpar := 0

    for i := 0; i < len(num); i++ {
        digito := int(num[i] - '0')

        if i%2 == 0 {
            somaPar += digito
        } else {
            somaImpar += digito
        }
    }

    dif := somaPar - somaImpar

    if dif < 0 {
        dif = -dif
    }

    if dif%11 == 0 {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}