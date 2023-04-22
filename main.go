package main

import (
	"fmt"
	"recipeFinder/cmd"
)

func main() {
	var ingrediens string
	var numberOfRecipes int
	ingrediens, numberOfRecipes = cmd.Execute()
	fmt.Println(ingrediens)
	fmt.Println(numberOfRecipes)

}
