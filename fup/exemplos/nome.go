package main
import "fmt"
func main() {

var nome string
var idade int
fmt.Println("Digite nome:")
fmt.Scan(&nome)
fmt.Println("Digite a idade:")
fmt.Scan(&idade)
fmt.Printf("Seu nome é %v e sua idade é %v\n", nome, idade)

}