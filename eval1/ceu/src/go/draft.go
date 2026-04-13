package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Print("[ ")
    for i := 0; i < 9; i++ {
        if i != n {
            fmt.Println(i)

            for a := i + 1; a < 9; a++ {
                
            }
        }
    }
    fmt.Println(" ]")
}
