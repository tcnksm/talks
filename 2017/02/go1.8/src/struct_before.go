package main

// START OMIT
type Person struct {
	Name string
}

type person struct {
	Name string `json:"name"`
}

func main() {
	p := person{
		Name: "tcnksm",
	}

	P := Person{
		Name: p.Name, // 😭
	}
}

// END OMIT
