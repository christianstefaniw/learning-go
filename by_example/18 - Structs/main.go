package main

import "fmt"

/*
create a TYPE called PERSON based on a struct of NAME string and AGE int
"type" is used to create a new type
 */
type person struct {
	name string
	age int
}

func newPerson(name string) *person{
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{name: "Bob", age: 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))


	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	/*
	You can also use dots with struct pointers - the pointers are automatically dereferenced.
	 */
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
