package main
import (
    "fmt"
    "bufio"
    "os"
)

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
    scanner := bufio.NewScanner(os.Stdin)
    var n int
    fmt.Scan(&n)
    
    alunos := make([]Aluno, n)
    for i := range alunos {
        scanner.Scan()
        alunos[i].nome = scanner.Text()
          scanner.Scan()
    fmt.Sscanf(scanner.Text(),
        "%f %f %f",
        &alunos[i].n1,
        &alunos[i].n2,
        &alunos[i].n3)

    alunos[i].media = mediaNotas(
        alunos[i].n1,
        alunos[i].n2,
        alunos[i].n3,
    )
    }
     ordenarAluno(alunos)
    for i := 0;  i < len(alunos); i++ {
        fmt.Printf("%d: %s\n   Media: %.2f\n   N1: %.2f, N2: %.2f, N3: %.2f\n", i, alunos[i].nome, alunos[i].media, alunos[i].n1, alunos[i].n2, alunos[i].n3)
    }
}

    