package main

import "fmt"

func main() {
	var valor float64
	var parcelas int
	fmt.Scan(&valor, &parcelas)
	var juros float64

	if parcelas == 1 {
		juros = 0
	} else {
		juros = float64(parcelas-1) * 0.05
	}

	total := valor * (1 + juros)
	valorParcela := total / float64(parcelas)

	fmt.Printf("%.2f\n", valorParcela)
	fmt.Printf("%.2f\n", total)
}