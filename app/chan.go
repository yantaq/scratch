package app

import "log"

// WriteToC writes x to chan c
func WriteToC(c chan int, x int) {
	log.Println("first: ", x)
	c <- x
	close(c)
	log.Println("second: ", x)
}
