package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    anterior := 0
    fmt.Scan(&anterior)

    var x int
    parkour := 0 
    for i := 1; i < n; i++{
        fmt.Scan(&x)
        if x-anterior > 1 || x - anterior < -1 {
            anterior = x
            parkour++
        } else {
            anterior = x
            continue
        }
    }
    fmt.Println(parkour)
}
