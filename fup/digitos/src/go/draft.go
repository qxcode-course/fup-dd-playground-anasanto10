package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

    contador := 0

    if b == 0{
        if a == 0 {
            fmt.Println("1")
        } else {
            fmt.Println("0")
        }
        return
    }
    
    for b > 0 {
        ult := b%10

        if ult == a {
            contador ++
        }
        b = b/10
    }
    fmt.Println(contador)
}
