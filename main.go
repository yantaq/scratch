package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/yantaq/scratch/app"
	foo "github.com/yantaq/scratch/pkg"
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
		app.Clone("https://github.com", "yantaq/scratch", "scratch")
		fmt.Println("git clone ...done.")
	case "pkg":
		fmt.Println("pkg example: ", foo.D())
	case "open":
		app.Open("")
	case "chan":
		c := make(chan int)
		go app.WriteToC(c, 20)
		time.Sleep(1 * time.Second)
		fmt.Println("Read: ", <-c)
		time.Sleep(1 * time.Second)
		_, ok := <-c
		if ok {
			fmt.Println("Channel is open!")
		} else {
			fmt.Println("Channel is closed!")
		}
	default:
		flag.PrintDefaults()
	}
}
