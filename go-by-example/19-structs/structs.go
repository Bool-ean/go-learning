package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(n string) *person{
	p := person{name: n}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "alice", age: 30})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person {name: "sean", age: 50}
	fmt.Println(s.name)

	sp := s 
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct{
		name string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
	
}
