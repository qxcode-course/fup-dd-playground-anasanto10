package main
import "fmt"
func main() {
    var N, X, Y, S int
    var C string
    fmt.Scan(&N, &X, &Y, &C, &S)

    if C == "R" {
        X = (X + S) % N
    } else if C == "L" {
        X = (X - S % N + N) % N
    } else if C == "D" {
        Y = (Y + S) % N
    } else if C == "U"{
        Y = (Y - S % N + N) % N
    }

    fmt.Println(X, Y)
}
