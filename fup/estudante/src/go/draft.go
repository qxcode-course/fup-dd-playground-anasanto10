package main
import "fmt"

type Aluno struct {
    nome string
    n1 float64
    n2 float64
    n3 float64
    media float64
}

func mediaNotas (n1, n2, n3 float64) float64 {
    media := (n1 + n2 + n3) / 3
    return media
}

func ordenarAluno (alunos []Aluno) {

    for i := 0; i < len(alunos) - 1; i++ {
        for j := 0; j < len(alunos) - 1; j++ {
            if alunos[j].media < alunos[j + 1].media {
                alunos[j], alunos[j + 1] = alunos[j + 1], alunos[j]
            }
        } 
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    alunos := make([]Aluno, n)
    for i := range alunos {
        fmt.Scan(&alunos[i].nome, &alunos[i].n1, &alunos[i].n2, &alunos[i].n3)
    }
}
    