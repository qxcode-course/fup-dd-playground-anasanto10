package main
import "fmt"
func main() {
    var valor, c1, c2 int
    fmt.Scan(&valor, &c1, &c2)

    dist1 := abs(valor - c1)
    dist2 := abs(valor - c2)

    if dist1 < dist2 {
        fmt.Println("primeiro")
    } else if dist2 < dist1{
        fmt.Println("segundo")
    } else {
        fmt.Println("empate")
    }
}

func abs(x int) int {
    if x < 0 {
    return - x
 }
  return x
}