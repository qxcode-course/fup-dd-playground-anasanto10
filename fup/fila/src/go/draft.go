package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

   impares := []int {}
   pares := []int {}

   for i:= 0; i < N; i++ {
    var num int
    fmt.Scan(&num)

    if num%2 != 0 {
        impares = append(impares, num)
    } else {
        pares = append(pares, num)
    }
   }

   fmt.Print("[ ")
   for _, valor := range impares {
    fmt.Print(valor, " ")
   }
   fmt.Println("]")

   fmt.Print("[ ")
   for _, valor := range pares {
    fmt.Print(valor, " ")
   }
   fmt.Println("]")
}
