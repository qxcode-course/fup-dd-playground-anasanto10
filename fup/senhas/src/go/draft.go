package main
import "fmt"
func main() {
    var s, q int
    fmt.Scan(& s, &q)
    var digitos string
    fmt.Scan(&digitos)
    var senha string
    fmt.Scan(&senha)

    vetorSenha := []byte(senha)

    for k := 0; k < q; k++ {

        for i := s - 1; i >= 0; i--{

            indice := 0 

            for j := 0; j < len(digitos); j++{
                if vetorSenha[i] == digitos[j] {
                    indice = j
                    break
                }
            }

            if indice < len(digitos)-1 {
                vetorSenha[i] = digitos[indice+1]
                break
            } else {
                vetorSenha[i] = digitos[0]
            }
        }

        fmt.Println(string(vetorSenha))
    }
}