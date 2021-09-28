package main

import (
	"fmt"

	"github.com/yantaq/scratch/app"
)

func main() {
	y := app.GCD(2, 4)
	fmt.Printf("gcd(%d, %d) => %d", 2, 4, y)
}
