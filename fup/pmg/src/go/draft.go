package main
import "fmt"
func main() {
	var s int
	fmt.Scan(&s)

	arr := make([]float64, s)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	var soma float64 = 0
	for i := 0; i < len(arr); i++ {
		soma += arr[i]
	}

	media := float64(soma) / float64(len(arr))
	fmt.Printf("%.2f\n", media)

	for i := 0; i < s; i++ {
		if arr[i] < media {
			fmt.Print("P")
		} else if arr[i] > media {
			fmt.Print("G")
		} else {
			fmt.Print("M")
		}

		if i != s-1 {
			fmt.Print(" ")
		}
	}
    fmt.Println()
}
