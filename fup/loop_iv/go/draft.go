package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    fmt.Print("[ ")
    if A > B {
        for i := A; i > B; i-- {
            if i != A {
                fmt.Print(" ")
            }
            fmt.Print(i)
        }
    } else if A < B {
        for i := A; i < B; i++ {
            if i != A {
                fmt.Print(" ")
            }
            fmt.Print(i)
        }
    }
    fmt.Println(" ]")
}
