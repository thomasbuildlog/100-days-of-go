package main

import "fmt"

const (
	Name = "Thomas"
	Goal = "Junior Go Developer in California"
	Days = 100
)

func main() {
	var age int = 23
	language := "Go"

	fmt.Printf("Hey, I'm %s, %d years old.\n", Name, age)
	fmt.Printf("Day 1/%d → learning %s\n", Days, language)
	fmt.Printf("Goal: %s\n", Goal)
}
