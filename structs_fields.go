package main

import (
	"fmt"
	"math/rand"
)

type Person struct {
	name   string
	age    int
	height float32
	next   *Person
}

func main() {
	var head *PersonPerson
	var tail *Person
	var a = 5
	for a > 0 {
		var ms Person
		general(&ms)

		if head == nil && tail == nil {
			head = &ms
			tail = &ms
		}

		tail.next = &ms
		tail = &ms
		a -= 1
	}

	for i := head; i != nil; i = i.next {
		fmt.Printf("The name is: %s\n", i.name)
		fmt.Printf("The age is: %d\n", i.age)
		fmt.Printf("The height is: %f\n", i.height)
		fmt.Println(*i)
	}
}

func general(p *Person) {
	p.age = rand.Intn(100)
	p.name = "name"
	p.height = rand.Float32() + 1.0
}
