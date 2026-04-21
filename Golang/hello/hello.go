package main

import (
	"fmt"
	"time"
)
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println("The time is", time.Now())
}
