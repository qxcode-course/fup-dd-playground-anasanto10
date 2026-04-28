package main
import  "fmt"

func main() {
    nums := 5
    arr := make([]int, nums)
     for i := 0; i < nums; i++{
        fmt.Scan(&arr[i])
     }
     a := arr[0]
     for _, inteiro := range arr {
        if inteiro < a {
            a = inteiro
        }

     }
    fmt.Println(a)

}
