package main
import "fmt"

func ehVogal(c rune) bool {
    return c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U' 
}

func contarVogais(nome string) int {
    cont := 0
    for _, c := range nome {
        if ehVogal(c) {
            cont += 1
        }
    }
    return cont
}

func main() {
    var nome1, nome2 string 
    fmt.Scan(&nome1, &nome2)
    var pontos int
    if nome1[0] == nome2[0] {
        pontos += 20
    }

    if len(nome1) == len(nome2) {
        pontos += 30
    }
     if contarVogais(nome1) == contarVogais(nome2) {
        pontos += 30
     }
     ult1 := rune(nome1[len(nome1) - 1])
     ult2 := rune(nome2[len(nome2) - 1])

     if ehVogal(ult1) == ehVogal(ult2) {
		pontos += 20
	} else {
		pontos -= 10
	}

     if pontos < 0 {
        pontos = 0
     }
     fmt.Printf("As chances do crush te dar bola sao: %d%%!\n", pontos)
}