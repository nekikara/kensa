package main

import (
	"fmt"
	"kensa/kensa"
)

func main() {
	spec := kensa.NewSpec()
	files := kensa.ListUpData()
	for _, file := range files {
		fmt.Print("A: " + spec.ExecuteScript(file))
		fmt.Print("B: " + kensa.ShowAnswer(file) + "\n\n")
	}
}
