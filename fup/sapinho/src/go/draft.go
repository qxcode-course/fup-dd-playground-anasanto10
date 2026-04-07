package main
import "fmt"
func main() {
    var P, S, E int
    fmt.Scan(&P, &S, &E)
    var next int
    pos := 0

    for {
      next = pos + S
      if next >= P {
        fmt.Printf("%d saiu\n", pos)
        break
      }

    fmt.Printf("%d %d\n", pos, next)
    pos = next - E
}
}