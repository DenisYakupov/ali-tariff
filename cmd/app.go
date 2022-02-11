package main

import "fmt"

type Tets struct {
	A int
	B string
}

func (t Tets) structs() {
	fmt.Println(t.A)
}
func main() {
	a := Tets{
		10,
		"test",
	}
	a.structs()
}
