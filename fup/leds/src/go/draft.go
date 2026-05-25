package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
    arr := make([]string, N)

    for i := range arr {
        fmt.Scan(&arr[i])
    }

    for i := 0; i < N; i++ {

        total := 0
        numero := arr[i]

        for j := 0; j <len(numero); j++ {
            switch numero[j] {

            case '0':
                total += 6

            case '1':
                total += 2

            case '2':
                total += 5

            case '3':
                total += 5

            case '4':
                total += 4

            case '5':
                total += 5

            case '6':
                total += 6

            case '7':
                total += 3

            case '8':
                total += 7

            case '9':
                total += 6
            }
        }
       fmt.Println(total, "leds")
    }
}

