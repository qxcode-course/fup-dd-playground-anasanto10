package main
import "fmt"
func main() {
    var a int16
    fmt.Scan(&a)
    if (a == 3 || a == 5){
    fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
