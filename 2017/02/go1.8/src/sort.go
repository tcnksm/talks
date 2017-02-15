package main

import "sort"

// START OMIT
type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age }) // 🙆
}

// END OMIT
