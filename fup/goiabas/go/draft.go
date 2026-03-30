package main
import "fmt"
func main() {
    var c, banana, goiaba, manga int16
    fmt.Scan(&c, &banana, &goiaba, &manga)

    total := banana + goiaba + manga
    viagens := total / c

    if (total%c == 0){
        fmt.Println(viagens)
    } else {
        viagens = viagens + 1
        fmt.Println(viagens)
    }
}
