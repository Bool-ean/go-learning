package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3{
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("direct")

	go func(msg string){
		fmt.Println(msg)
	}("going")
	
	go f("goroutine")


	time.Sleep(time.Second)
	fmt.Println("done")
}
