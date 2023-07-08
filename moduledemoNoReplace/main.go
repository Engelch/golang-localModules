package main

import (
	"fmt"
	"mypackage" // Importing mypackage packages under the same project
)

func main() {
	mypackage.New()
	fmt.Println("main")
}
