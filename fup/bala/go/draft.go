package main

import (
	"fmt"
	"math"
)
func main() {
    var x1, x2, y1, y2, dist float64
    fmt.Scan(&x1, &x2, &y1, &y2)
    dist = math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
    fmt.Printf("%.2f\n", dist)
}
