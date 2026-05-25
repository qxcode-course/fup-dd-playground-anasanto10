package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)

    for i := range arr {
        fmt.Scan(&arr[i])
    }
     casais := 0 

    for i := 0; i < n; i++ {
        if arr[i] > 0 {
            
            for j := 0; j < n; j++{

                if arr[j] == -arr[i] {
                casais++
                arr[j] = 0
                arr[i] = 0
                break
                }
            }
        }
    }
    fmt.Println(casais)
}
