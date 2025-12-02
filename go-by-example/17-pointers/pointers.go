package main

import "fmt"

func zeroval(ival int){
	ival = 0
}

func zerptr(iptr *int){
	*iptr = 0
}

func main() {
	i :=1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zerptr(&i)
	fmt.Println("zerptr:", i)

	fmt.Println("pointer:", &i)
}
