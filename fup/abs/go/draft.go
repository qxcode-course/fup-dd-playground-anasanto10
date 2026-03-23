package main
import "fmt"
func main() {
    var a, b, dif, abs int16
    fmt.Scan(&a, &b)
    dif = a - b
    if(dif < 0) {
     abs = dif * (-1)
     fmt.Println(abs)
    } else {
        fmt.Println(dif)
    }
}
