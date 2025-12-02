package main

import (
	"fmt"
	"time"
)

func main() {
	
	i := 2
	fmt.Print("Write ", i, " as ")
	
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch{
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("It's after noon")
	}


	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("im a bool")
		case int:
			fmt.Println("im an int")
		default:
			fmt.Println("dont know what type", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hi")

}
