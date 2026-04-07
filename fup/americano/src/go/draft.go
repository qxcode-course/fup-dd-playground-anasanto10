package main
import "fmt"
func main() {
    var n1, n2, n3, n4 int
    fmt.Scan(&n1, &n2, &n3, &n4)

    soma := n1 + n2 + n3 + n4
    resto := soma%4

    if soma == 0{
        fmt.Println("nenhum")
    } else if resto == 1 {
        fmt.Println("jog1")
    } else if resto == 2 {
        fmt.Println("jog2")
    } else if resto == 3 {
        fmt.Println("jog3")
    } else {
        fmt.Println("jog4")
    }
}
