package main
import "fmt"
func main() {
   //var qtd int 
   //fmt.Scan(&qtd)
   //var idades[]int = make([]int, qtd)
   //fmt.Println(idades)

   //arr := []int {9, 8, 7, 6, 5, 4}
   //fmt.Print("[ ")
   //for i:=0; i < len(arr); i++{
   // fmt.Printf("%d ", arr[i])
   //}
   //fmt.Println("]")

   //arr := []int {9, 8, 7, 6, 5, 4, 3, 2}
   //fmt.Print("[ ")
   //for pos, _ := range arr {
   // fmt.Printf("%d ", pos)
   //}
   //fmt.Print("]\n")

    func mostrar_vetor (arr []int, sep string) {
    arr := []int {9, 8, 7, 6, 5, 4, 3, 2}
    fmt.Print("[ ")
    for i, valor := range arr {
     if i != 0 {
        fmt.Print(sep)
     }
     fmt.Printf("%v", valor)
    }
    fmt.Print("]\n")
}

    func main() {
        var qtd int
        fmt.Scan(&qtd)
        var arr []int = make([]int, qtd)
        for i := range arr {
            fmt.Scan(&arr[i])
        }
        mostrar_vetor(arr, " ")
    }

}





