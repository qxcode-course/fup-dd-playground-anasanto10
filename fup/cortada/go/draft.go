package main
import "fmt"
func main() {
    var B, T int
    fmt.Scan(&B, &T)

    sum := B + T

    if sum > 160 {
        fmt.Println(1)
    } else if sum < 160 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}