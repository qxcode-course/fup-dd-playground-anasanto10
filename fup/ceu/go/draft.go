package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    fmt.Print("[ ")
    for i := 0; i <= 9; i++ {
          if i != N {
            fmt.Print(i)

            for a := i + 1; a <= 9; a++ {
                if a != N {
                    fmt.Print(" ", a)
                }
            }
            break
    }
  }
        if N == 10 {
        fmt.Println(" ]")
    } else {
        fmt.Println(" ceu ]")
    }
}
