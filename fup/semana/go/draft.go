package main
import "fmt"
func main() {
    var dia, hora int16
    fmt.Scan(&dia, &hora)

    if (dia >= 2 && dia <=6 && ((hora >=8 && hora <=11) || (hora >= 14 && hora <= 17))) ||
    (dia == 7 && (hora >=8 && hora <= 11)) {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
