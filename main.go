package main

import (
	"flag"
	"fmt"

	"github.com/yantaq/scratch/app"
)

func main() {

	var ops = flag.String("ops", "", "flag required \"-ops string\" i.e: -ops gcd")
	flag.Parse()

	switch *ops {
	case "gcd":
		y := app.GCD(2, 4)
		fmt.Printf("gcd(%d, %d) => %d", 2, 4, y)
	case "nfib":
		y := app.Fib(4)
		fmt.Printf("nfib(%d) => %d", 4, y)
	case "clone":
		app.Clone()
		fmt.Println("git clone ...done.")
	default:
		flag.PrintDefaults()
	}
}
