package main
import "fmt"
func main() {
    var nome string
    var idd int
    fmt.Scan(&nome, &idd)
    if (idd < 12){
        fmt.Println(nome,"eh crianca")
    } else if (idd < 18) {
        fmt.Println(nome,"eh jovem")
    } else if (idd < 65) {
        fmt.Println(nome,"eh adulto")
    } else if (idd < 1000) {
        fmt.Println(nome,"eh idoso")
    } else {
        fmt.Println(nome,"eh mumia")
    }
}
