package main
import "fmt"

func mmc(a, b int) int{
    maior := a
    if b > maior {
        maior = b
    }
    for {
        if maior%a == 0 && maior%b == 0 {
            return maior
        }
        maior ++
    }
}
func main() {
    var N, M int
    fmt.Scan(&N, &M)

    resultado := mmc(N, M)
    fmt.Println(resultado)
}