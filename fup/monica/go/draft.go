package main
import "fmt"
func main() {
    var M, A, B int
    fmt.Scan(&M, &A, &B)

    C := M - (A + B)

    if (A > B && A > C){
        fmt.Println(A)
    } else if (B > A && B > C){
        fmt.Println(B)
    } else {
        fmt.Println(C)
    }
}
