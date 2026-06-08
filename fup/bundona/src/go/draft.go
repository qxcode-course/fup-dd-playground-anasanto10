package main
import "fmt"
func main() {
	var h, m, d int
	var s string

	fmt.Scan(&h)
	fmt.Scan(&m)
	fmt.Scan(&s)
	fmt.Scan(&d)

	pos := h*6 + m/10

	if s == "H" {
		pos += d
	} else {
		pos -= d
	}

	pos = ((pos % 72) + 72) % 72

	h = pos / 6
	m = (pos % 6) * 10

	fmt.Printf("%02d %02d\n", h, m)
}