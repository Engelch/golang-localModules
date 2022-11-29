package main

import (
	"fmt"
	"moduledemoNoReplace/mypackage" // Importing mypackage packages under the same project
)

func main() {
	mypackage.New()
	fmt.Println("main")
}
