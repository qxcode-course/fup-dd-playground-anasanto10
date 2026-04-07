package main
import "fmt"
func main() {
    var v1, v2, v3 int
    fmt.Scan(&v1, &v2, &v3)

    if v1 == v2 && v1 == v3 {
        fmt.Println("empate")
    } else if (v1 == 1 && v2 == 0 && v3 == 0) || (v1 == 0 && v2 == 1 && v3 == 1) {
        fmt.Println("jog1")
    } else if (v1 == 1 && v2 == 0 && v3 == 1) || (v1 == 0 && v2 == 1 && v3 == 0){
        fmt.Println("jog2")
    } else {
        fmt.Println("jog3")
    }
}
