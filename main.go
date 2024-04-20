package main

import (
	"fmt"
)

var names = []string{"nana", "Kwafo"}

func main() {
	names = append(names, "Mensah")
	fmt.Println(names[2])
}
