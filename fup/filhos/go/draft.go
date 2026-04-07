package main
import "fmt"
func main() {
    var idd1, quant int
    fmt.Scan(&idd1, &quant)

    iddV := idd1 + (quant * 2)
    for i := idd1; i < iddV; i += 2 {
        fmt.Println(i)
    }
}
