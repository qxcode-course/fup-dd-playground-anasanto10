package main
import "fmt"
func main() {
    var A, B int16
    fmt.Scan (&A, &B)
    soma := A + B
    sub := A - B
    mult := A * B
    div := float64(A) / float64(B)
    rest := A % B
    fmt.Println(soma)
    fmt.Println(sub)
    fmt.Println(mult)
    fmt.Printf("%.2f\n",div)
    fmt.Println(rest)
}
