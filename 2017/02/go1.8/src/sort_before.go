package main

import "sort"

type Person struct {
	Name string
	Age  int
}

// START OMIT
type byAge []Person

func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age } // 😭
func (a byAge) Len() int           { return len(a) }              // 😭
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }    // 😭

func main() {
	people := []Person{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Sort(byAge(people))
}

// END OMIT
