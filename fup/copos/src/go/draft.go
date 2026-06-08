package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    largura := 2*n - 1 

    for i := 1; i <= n; i++ {

        for j := 0; j < largura; j++ {

            if j >= n-i && (j-(n-i))%2 == 0 && (j-(n-i))/2 < i {
                fmt.Print(n)
            } else {
                fmt.Print(".")
            }
        }

        fmt.Println()
    }
}
