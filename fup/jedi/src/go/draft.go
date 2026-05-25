package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    metade := n/2

    somJedi := 0
    somSith := 0

    for i := 0; i < metade; i++ {
        somJedi += arr[i]
    }

    for i := metade; i < n; i++ {
        somSith += arr[i]
    }

    if somJedi < somSith {
        fmt.Println("Sith")
    } else if somSith < somJedi {
        fmt.Println("Jedi")
    } else {
        fmt.Println("Empate")
    }
}
