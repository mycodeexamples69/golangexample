package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("The time is", time.Now())
	fmt.Printf("\n")
	fmt.Println("My number is", rand.Intn(10))
	// print name
	name := "Josh Smith"

	fmt.Printf("Hello, %s!\n", name)

	idx := 5

	if len(name) >= idx {
		fmt.Printf("Name is longer than 10\n")
		fmt.Printf("Name length is %d", len(name))
	}

	// while loop
	for idx >= 10 {
		fmt.Printf("Loop Count %d\n", idx)
	}

	// for loop

	count := 0
	for i := 0; i < 10; i++ {
		count += i
	}
	fmt.Printf("Count %d\n", count)
}
