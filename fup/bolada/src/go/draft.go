package main
import "fmt"
func main() {
    var saque string
    var forca, T int
    fmt.Scan(&saque, &forca)

    if saque == "b" {
        T = 20
    } else {
        T = 18
    }

    P := ((forca * T) - 80) / 10

    if P < 150 {
        fmt.Println("Fraco, nem passou")
    } else if P >= 150 && P  < 180 {
        fmt.Println("Perfeito")
    } else if P >= 180 && P <= 210 {
        fmt.Println("Satisfeito")
    } else {
        fmt.Println("Muito forte, bola fora")
    }
}
