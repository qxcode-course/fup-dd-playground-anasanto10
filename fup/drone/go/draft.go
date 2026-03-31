package main
import "fmt"
func main() {
    var A, B, C, H, L int
    fmt.Scan(&A, &B, &C, &H, &L)

    if (A <= H && B <= L) || (A <= L && B <= H) || (A <= H && C <= L) || (A <= L && C <= H) || (B <= H && C <= L) || (B <= L && C <= H){
        fmt.Println("S")
    } else{
        fmt.Println("N")
    }
}
