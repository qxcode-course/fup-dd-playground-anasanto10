package main
import "fmt"
func main() {
    var N int
    var pe string
    fmt.Scan(&N, &pe)

     fmt.Print("[ ")
    for i := 0; i <= 9; i++ {
          if i != N {
            fmt.Print(i, pe)

            if pe == "e" {
                pe = "d"
            } else {
                pe = "e"
            }

            for a := i + 1; a <= 9; a++ {
                if a != N {
                    fmt.Print(" ", a, pe)
                    if pe == "e" {
                        pe = "d"
                    } else {
                        pe = "e"
                    }
                }
            }
            break
        }
    }

    if N != 10 {
        fmt.Print(" ceu")
    }
    fmt.Println(" ]")
}


