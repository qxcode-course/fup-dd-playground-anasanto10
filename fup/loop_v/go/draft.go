package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    i:= A
    fmt.Print("[ ")
    for{
        if i == B{
            break
        }
        if i%2 == 0{
            i++
            continue
        }
        fmt.Print(i, " ")
        i++
    }
    fmt.Println("]")
}
