package main
import "fmt"

type Pessoa struct {
    nome string
    idade int
    sexo string
}

func maisIdosa (pessoas []Pessoa) {
   maisVelha := make([]Pessoa, 0)

   for i := 0; i < len(pessoas) - 1; i++ {
    if pessoas[i].sexo == "f" {
        maisVelha = pessoas[i]
    }
   }

   for i := 1; i < len(maisVelha); i++ {
    if maisVelha < pessoas[i] {
        maisVelha = pessoas [i]
    }
   }
}

func main() {
    var n int
    fmt.Scan(&n)

    pessoas := make([]Pessoa, n)
    for i := range pessoas {
        fmt.Scan(&pessoas[i].nome, &pessoas[i].idade, &pessoas[i].sexo)
    }


}