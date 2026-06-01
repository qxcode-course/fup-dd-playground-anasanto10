package main
import "fmt"
func main() {
    var tam int
    fmt.Scan(&tam)

    arr := make([]int, tam)
    for i := range arr {
        fmt.Scan(&arr[i])
    }

    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); i++ {

            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }

    fmt.Println(arr)
}