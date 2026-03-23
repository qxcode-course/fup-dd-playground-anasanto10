package main
import "fmt"

func main() {
    var op string
    var num float64
    fmt.Scan(&op, &num)

    inteiro := int(num)
    decimal := num - float64(inteiro)

     if (op == "f") {
        fmt.Println(inteiro)
    } else if (op == "c") {
        if (decimal > 0) {
            fmt.Println(inteiro + 1)
        } else {
            fmt.Println(inteiro)
        }
    } else if (op == "r") {
         if (decimal >= 0.5) {
            fmt.Println(inteiro + 1)
        } else {
            fmt.Println(inteiro)
        }
    }
}