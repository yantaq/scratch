package main

import (
	"flag"
	"fmt"

	"github.com/yantaq/scratch/app"
)

func main() {

	var ops = flag.String("ops", "", "compute operations based on -ops flag value")
	flag.Parse()

	switch *ops {
	case "gcd":
		y := app.GCD(2, 4)
		fmt.Printf("gcd(%d, %d) => %d", 2, 4, y)
	case "nfib":
		y := app.Fib(4)
		fmt.Printf("nfib(%d) => %d", 4, y)
	default:
		fmt.Println("missing operations flag: --gcd, --nfib ...")
	}
}
