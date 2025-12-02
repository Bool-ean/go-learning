package main

import(
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 5000000000

	const d = 3e20

	fmt.Println(d)
	fmt.Println(d/n)

	fmt.Println(int64(d/n))

	fmt.Println(math.Sin(n))
}
